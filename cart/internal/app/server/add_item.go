package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/domain"
	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type AddItemRequest struct {
	Count  domain.Count `json:"count" validate:"min=1"`
	SKU    domain.Sku   `validate:"min=1"`
	UserId domain.UID   `validate:"min=1"`
}

const PathAddItem = "POST /user/{user_id}/cart/{sku_id}"

func (s *CartServer) ExtractAddItemRequest(r *http.Request) (addItemRequest *AddItemRequest, err error) {
	addItemRequest = &AddItemRequest{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, addItemRequest)
	if err != nil {
		return nil, err
	}

	rawID := r.PathValue("user_id")
	addItemRequest.UserId, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, err
	}

	rawID = r.PathValue("sku_id")
	addItemRequest.SKU, err = strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		return nil, err
	}

	// Валидация запроса
	_, err = s.validator.Validate(*addItemRequest)
	if err != nil {
		return nil, err
	}

	return addItemRequest, nil
}

// "POST /user/{user_id}/cart/{sku_id}"
func (s *CartServer) AddItem(w http.ResponseWriter, r *http.Request) {

	addItemRequest, err := s.ExtractAddItemRequest(r)

	if err != nil {
		serverErr := ServerError{
			Path:   PathAddItem,
			Status: http.StatusBadRequest,
			Text:   err.Error(),
		}
		writeErrorResponse(w, &serverErr)
		return
	}

	_, err = s.cartService.AddItem(r.Context(), model.CartItem{
		SKU:    addItemRequest.SKU,
		UserId: addItemRequest.UserId,
		Count:  addItemRequest.Count,
	})
	if err != nil {
		serverErr := ServerError{
			Path:   PathAddItem,
			Status: http.StatusInternalServerError,
			Text:   err.Error(),
		}

		if serverErr.Text == "sku does not exist" || serverErr.Text == "rpc error: code = FailedPrecondition desc = wrong sku" {
			serverErr.Status = http.StatusPreconditionFailed
		}

		writeErrorResponse(w, &serverErr)
		return
	}

	fmt.Fprint(w, "{}")

}
