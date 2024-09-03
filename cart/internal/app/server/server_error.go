package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ServerError struct {
	Status int
	Text   string
	Path   string
}

func (e *ServerError) Error() string {
	return e.Path + ": " + e.Text
}

type ErrorResponse struct {
	Message string `json: "message"`
}

func writeErrorResponse(w http.ResponseWriter, err *ServerError) {
	w.WriteHeader(err.Status)
	w.Header().Set("Content-Type", "application/json")

	jsonBytes, jsonErr := json.Marshal(ErrorResponse{
		Message: err.Error(),
	})

	if jsonErr != nil {
		panic(jsonErr)
	}

	_, errOut := fmt.Fprint(w, string(jsonBytes))
	if errOut != nil {
		log.Printf("errOut.Error()")
		return
	}
}
