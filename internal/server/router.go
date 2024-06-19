package server

import (
	"net/http"

	"github.com/NicholeMattera/OutClimb-Event-Manager/internal/utils"
)

func SetupRouter() {
	mux := http.NewServeMux()
	db := utils.Database()

	NewPingHandler(mux)
	NewCheckHandler(mux, db)
	NewRegistrationHandler(mux, db)

	http.ListenAndServe(":80", mux)
}
