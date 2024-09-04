package server

import (
	"context"
	"net/http"
	"strconv"
)

// "DELETE /user/{user_id}/cart"
func (s *CartServer) DeleteCart(w http.ResponseWriter, r *http.Request) {
	serverErr := ServerError{
		Path: "DELETE /user/{user_id}/cart",
	}

	rawID := r.PathValue("user_id")
	userId, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		serverErr.Status = http.StatusBadRequest
		serverErr.Text = "can't parse user_id"
		writeErrorResponse(w, &serverErr)
		return
	}

	err = s.cartService.DeleteCart(context.Background(), userId)
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
