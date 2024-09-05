package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type GetItemsRequest struct {
	UserId model.UID `validate:"min=1"`
}

// "GET /user/{user_id}/cart"
func (s *CartServer) GetItems(w http.ResponseWriter, r *http.Request) {
	getItemsRequest := GetItemsRequest{}
	serverErr := ServerError{
		Path: "GET /user/{user_id}/cart",
	}

	rawID := r.PathValue("user_id")
	var err error
	getItemsRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse user_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	// Валидация запроса
	if _, err := s.validator.Validate(getItemsRequest); err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	result, err := s.cartService.GetItems(context.Background(), getItemsRequest.UserId)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	rawResponse, err := json.Marshal(result)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	if len(result.Items) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Fprint(w, string(rawResponse))
	}
}
