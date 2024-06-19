package server

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong\n"))
}

func check(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://outclimb.gay/event-registration-form", http.StatusTemporaryRedirect)
}

func registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success":true}`))
}

func SetupRouter() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)
	mux.HandleFunc("GET /check", check)
	mux.HandleFunc("POST /registration", registration)

	http.ListenAndServe(":80", mux)
}
