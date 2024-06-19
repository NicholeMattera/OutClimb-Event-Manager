package server

import (
	"net/http"
)

type PingHandler struct{}

func NewPingHandler(mux *http.ServeMux) {
	pingHandler := &PingHandler{}

	mux.HandleFunc("GET /ping", pingHandler.Ping())
}

func (*PingHandler) Ping() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong\n"))
	})
}
