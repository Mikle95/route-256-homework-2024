package server

import (
	"context"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type DeleteCartRequest struct {
	UserId model.UID `validate:"min=1"`
}

// "DELETE /user/{user_id}/cart"
func (s *CartServer) DeleteCart(w http.ResponseWriter, r *http.Request) {
	deleteCartRequest := DeleteCartRequest{}

	serverErr := ServerError{
		Path: "DELETE /user/{user_id}/cart",
	}

	rawID := r.PathValue("user_id")
	var err error
	deleteCartRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse user_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	// Валидация запроса
	if _, err := s.validator.Validate(deleteCartRequest); err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.cartService.DeleteCart(context.Background(), deleteCartRequest.UserId)
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
