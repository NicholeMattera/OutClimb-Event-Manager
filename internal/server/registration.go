package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/NicholeMattera/OutClimb-Event-Manager/internal/model"
)

type registrationRequest struct {
	Event  string `json:"event"`
	Secret string `json:"secret"`
}

type RegistrationHandler struct {
	db *sql.DB
}

func NewRegistrationHandler(mux *http.ServeMux, db *sql.DB) {
	registrationHandler := &RegistrationHandler{
		db: db,
	}

	mux.HandleFunc("POST /registration", registrationHandler.Register())
}

func (rh *RegistrationHandler) Register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Decode the request body
		var body registrationRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"success":false}`))
			return
		}

		// Check the secret
		secret, secretExists := os.LookupEnv("OUTCLIMB_REGISTRATION_SECRET")
		if secretExists && body.Secret != secret {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"success":false}`))
			return
		}

		event, _ := model.GetEvent(rh.db, body.Event)
		if event.Id == 0 {
			model.CreateEvent(rh.db, body.Event)
		} else {
			event.NumberOfRegistrations++
			model.UpdateEvent(rh.db, &event)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success":true}`))
	})
}
