package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("internal/web/static/"))))

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/status", StatusHandler).Methods("GET")
	r.HandleFunc("/api/audit", AuditHandler).Methods("GET")
	r.HandleFunc("/api/heal", ManualHealHandler).Methods("POST")

	return r
}
