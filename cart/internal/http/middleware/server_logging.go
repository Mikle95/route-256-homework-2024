package middleware

import (
	"log"
	"net/http"
)

type LogMux struct {
	h http.Handler
}

func NewLogMux(h http.Handler) http.Handler {
	return &LogMux{h: h}
}

func (m *LogMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: %v%v ContentLength:%v", r.Method, r.Host, r.URL, r.ContentLength)
	m.h.ServeHTTP(w, r)
}
