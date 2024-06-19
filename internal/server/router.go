package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/NicholeMattera/OutClimb-Event-Manager/internal/model"
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
		latestEvent, _ := model.GetEvent(db, "20240727-outdoor-climbing-he-mni-can-barn-bluff")

		if latestEvent.NumberOfRegistrations < 12 {
			http.Redirect(w, r, "https://outclimb.gay/event-registration-form", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "https://outclimb.gay/event-registration-filled", http.StatusTemporaryRedirect)
		}
	})
}

type registrationRequest struct {
	Event string `json:"event"`
}

func registrationHandler(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var body registrationRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"success":false}`))
			return
		}

		event, _ := model.GetEvent(db, body.Event)
		if event.Id == 0 {
			model.CreateEvent(db, body.Event)
		} else {
			event.NumberOfRegistrations++
			model.UpdateEvent(db, &event)
		}

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
