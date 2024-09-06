package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type GetItemsRequest struct {
	UserId model.UID `validate:"min=1"`
}

var PathGetItems = "GET /user/{user_id}/cart"

func (s *CartServer) ExtractGetItemsRequest(r *http.Request) (getItemsRequest *GetItemsRequest, err error) {
	getItemsRequest = &GetItemsRequest{}
	rawID := r.PathValue("user_id")
	getItemsRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, err
	}

	// Валидация запроса
	_, err = s.validator.Validate(*getItemsRequest)
	if err != nil {
		return nil, err
	}

	return getItemsRequest, nil
}

// "GET /user/{user_id}/cart"
func (s *CartServer) GetItems(w http.ResponseWriter, r *http.Request) {
	getItemsRequest, err := s.ExtractGetItemsRequest(r)
	if err != nil {
		serverErr := ServerError{
			Status: http.StatusBadRequest,
			Text:   err.Error(),
			Path:   PathGetItems,
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.SendGetItemsResponse(w, r, getItemsRequest)
	if err != nil {
		serverErr := ServerError{
			Status: http.StatusInternalServerError,
			Text:   err.Error(),
			Path:   PathGetItems,
		}
		writeErrorResponse(w, &serverErr)
		return
	}
}

func (s *CartServer) SendGetItemsResponse(w http.ResponseWriter, r *http.Request, getItemsRequest *GetItemsRequest) (err error) {
	result, err := s.cartService.GetItems(r.Context(), getItemsRequest.UserId)
	if err != nil {
		return err
	}

	rawResponse, err := json.Marshal(result)
	if err != nil {
		return err
	}

	if len(result.Items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{}")
	} else {
		fmt.Fprint(w, string(rawResponse))
	}
	return nil
}
