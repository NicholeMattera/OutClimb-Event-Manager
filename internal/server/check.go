package server

import (
	"database/sql"
	"net/http"

	"github.com/NicholeMattera/OutClimb-Event-Manager/internal/model"
)

type CheckHandler struct {
	db *sql.DB
}

func NewCheckHandler(mux *http.ServeMux, db *sql.DB) {
	checkHandler := &CheckHandler{
		db: db,
	}

	mux.HandleFunc("GET /check", checkHandler.Check())
}

func (ch *CheckHandler) Check() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		latestEvent, _ := model.GetEvent(ch.db, "20240810-outdoor-climbing-sugar-loaf-bluff")

		if latestEvent.NumberOfRegistrations < 12 {
			http.Redirect(w, r, "https://outclimb.gay/event-registration-form", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(w, r, "https://outclimb.gay/event-registration-filled", http.StatusTemporaryRedirect)
		}
	})
}
