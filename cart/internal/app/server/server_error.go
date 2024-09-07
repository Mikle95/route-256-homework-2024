package server

import (
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

func writeErrorResponse(w http.ResponseWriter, err *ServerError) {
	w.WriteHeader(err.Status)
	log.Println(err.Error())
	fmt.Fprint(w, "{}")
}
