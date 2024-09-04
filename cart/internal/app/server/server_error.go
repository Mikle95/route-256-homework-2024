package server

import (
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
	// w.Header().Set("Content-Type", "application/json")

	// jsonBytes, jsonErr := json.Marshal(Response{
	// 	Message: err.Error(),
	// })

	// if jsonErr != nil {
	// 	panic(jsonErr)
	// }

	// _, errOut := fmt.Fprint(w, string(jsonBytes))
	// if errOut != nil {
	// 	log.Printf("errOut.Error()")
	// 	return
	// }
}
