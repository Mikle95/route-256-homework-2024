package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// "GET /user/{user_id}/cart"
func (s *CartServer) GetItems(w http.ResponseWriter, r *http.Request) {
	serverErr := ServerError{
		Path: "GET /user/{user_id}/cart",
	}

	rawID := r.PathValue("user_id")
	userId, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse user_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	result, err := s.cartService.GetItems(context.Background(), userId)
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
