package middleware

import (
	"log"
	"net/http"
)

type LogMux struct {
	h http.HandlerFunc
}

func NewLogMux(h http.HandlerFunc) http.Handler {
	return &LogMux{h: h}
}

func (m *LogMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	m.h(w, r)
}
