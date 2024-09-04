package server

import (
	"context"
	"net/http"
	"strconv"
)

// "DELETE /user/{user_id}/cart/{sku_id}"
func (s *CartServer) DeleteItem(w http.ResponseWriter, r *http.Request) {
	serverErr := ServerError{
		Path: "DELETE /user/{user_id}/cart/{sku_id}",
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

	err = s.cartService.DeleteItem(context.Background(), userId, sku)
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
