package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"gitlab.ozon.dev/1mikle1/homework/cart/internal/pkg/cart/model"
)

type AddItemRequest struct {
	Count model.Count `json:"count"`
}

type AddItemResponse struct {
	SKU    model.Sku   `json:"sku"`
	Count  model.Count `json:"count"`
	UserId model.UID   `json:"user_id"`
}

// "POST /user/{user_id}/cart/{sku_id}"
func (s *CartServer) AddItem(w http.ResponseWriter, r *http.Request) {
	serverErr := ServerError{
		Path: "POST /user/{user_id}/cart/{sku_id}",
	}

	rawID := r.PathValue("user_id")
	userId, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse user_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	rawID = r.PathValue("sku_id")
	sku, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse sku_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	var addItemRequest AddItemRequest
	err = json.Unmarshal(body, &addItemRequest)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	_, err = s.cartService.AddItem(context.Background(), model.CartItem{
		SKU:    sku,
		UserId: userId,
		Count:  addItemRequest.Count,
	})
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = err.Error()
		writeErrorResponse(w, &serverErr)
		return
	}

	// rawResponse, err := json.Marshal(&AddItemResponse{
	// 	SKU:    item.SKU,
	// 	Count:  item.Count,
	// 	UserId: item.UserId,
	// })
	// if err != nil {
	// 	serverErr.Status = http.StatusBadRequest
	// 	serverErr.Text = err.Error()
	// 	writeErrorResponse(w, &serverErr)
	// 	return
	// }

	// fmt.Fprint(w, string(rawResponse))

}
