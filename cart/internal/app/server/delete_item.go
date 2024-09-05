package server

import (
	"context"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type DeleteItemRequest struct {
	SKU    model.Sku `validate:"min=1"`
	UserId model.UID `validate:"min=1"`
}

// "DELETE /user/{user_id}/cart/{sku_id}"
func (s *CartServer) DeleteItem(w http.ResponseWriter, r *http.Request) {
	deleteItemRequest := DeleteItemRequest{}

	serverErr := ServerError{
		Path: "DELETE /user/{user_id}/cart/{sku_id}",
	}

	rawID := r.PathValue("user_id")
	var err error
	deleteItemRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse user_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	rawID = r.PathValue("sku_id")
	deleteItemRequest.SKU, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse sku_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	// Валидация запроса
	if _, err := s.validator.Validate(deleteItemRequest); err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.cartService.DeleteItem(context.Background(), deleteItemRequest.UserId, deleteItemRequest.SKU)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	// rawResponse, err := json.Marshal(&Response{
	// 	Message: "Ok!",
	// })
	// if err != nil {
	// 	serverErr.Status = http.StatusBadRequest
	// 	serverErr.Text = err.Error()
	// 	writeErrorResponse(w, &serverErr)
	// 	return
	// }

	// fmt.Fprint(w, string(rawResponse))

}
