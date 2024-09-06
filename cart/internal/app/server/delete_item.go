package server

import (
	"fmt"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
)

type DeleteItemRequest struct {
	SKU    domain.Sku `validate:"min=1"`
	UserId domain.UID `validate:"min=1"`
}

func (s *CartServer) ExtractDeleteItemRequest(r *http.Request) (deleteItemRequest *DeleteItemRequest, err error) {
	deleteItemRequest = &DeleteItemRequest{}
	rawID := r.PathValue("user_id")
	deleteItemRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, err
	}

	rawID = r.PathValue("sku_id")
	deleteItemRequest.SKU, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, err
	}

	// Валидация запроса
	_, err = s.validator.Validate(*deleteItemRequest)
	if err != nil {
		return nil, err
	}

	return deleteItemRequest, nil
}

var PathDeleteItem = "DELETE /user/{user_id}/cart"

// "DELETE /user/{user_id}/cart/{sku_id}"
func (s *CartServer) DeleteItem(w http.ResponseWriter, r *http.Request) {
	deleteItemRequest, err := s.ExtractDeleteItemRequest(r)
	if err != nil {
		serverErr := ServerError{
			Status: http.StatusBadRequest,
			Text:   err.Error(),
			Path:   PathDeleteItem,
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.cartService.DeleteItem(r.Context(), deleteItemRequest.UserId, deleteItemRequest.SKU)
	if err != nil {
		serverErr := ServerError{
			Status: http.StatusInternalServerError,
			Text:   err.Error(),
			Path:   PathDeleteItem,
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "{}")
}
