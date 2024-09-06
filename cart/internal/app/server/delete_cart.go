package server

import (
	"fmt"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type DeleteCartRequest struct {
	UserId model.UID `validate:"min=1"`
}

var PathDeleteCart = "DELETE /user/{user_id}/cart"

func (s *CartServer) ExtractDeleteCartRequest(r *http.Request) (deleteCartRequest *DeleteCartRequest, err error) {
	deleteCartRequest = &DeleteCartRequest{}
	rawID := r.PathValue("user_id")
	deleteCartRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, err
	}

	// Валидация запроса
	_, err = s.validator.Validate(*deleteCartRequest)
	if err != nil {
		return nil, err
	}

	return deleteCartRequest, nil
}

// "DELETE /user/{user_id}/cart"
func (s *CartServer) DeleteCart(w http.ResponseWriter, r *http.Request) {

	deleteCartRequest, err := s.ExtractDeleteCartRequest(r)

	if err != nil {
		serverErr := ServerError{
			Status: http.StatusBadRequest,
			Text:   err.Error(),
			Path:   PathDeleteCart,
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.cartService.DeleteCart(r.Context(), deleteCartRequest.UserId)
	if err != nil {
		serverErr := ServerError{
			Path:   PathDeleteCart,
			Status: http.StatusInternalServerError,
			Text:   err.Error(),
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "{}")

}
