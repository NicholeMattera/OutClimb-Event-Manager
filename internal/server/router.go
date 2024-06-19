package server

import (
	"database/sql"
	"net/http"
)

func pingHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong\n"))
	})
}

func checkHandler(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://outclimb.gay/event-registration-form", http.StatusTemporaryRedirect)
	})
}

func registrationHandler(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true}`))
	})
}

func SetupRouter() {
	db := Database()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", pingHandler())
	mux.HandleFunc("GET /check", checkHandler(db))
	mux.HandleFunc("POST /registration", registrationHandler(db))

	http.ListenAndServe(":80", mux)
}
