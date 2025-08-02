package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", serveIndex).Methods("GET")
	return r
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Automation Smart Warehouse Platform"))
}
