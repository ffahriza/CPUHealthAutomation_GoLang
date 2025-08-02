package main

import (
	"log"
	"net/http"

	"github.com/ffahriza/automation-smart-warehouse/internal/audit"
	"github.com/ffahriza/automation-smart-warehouse/internal/config"
	"github.com/ffahriza/automation-smart-warehouse/internal/engine"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services and healing logic
	go engine.StartScheduler(cfg)

	// Setup HTTP router
	r := mux.NewRouter()

	// Serve dashboard at "/"
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "internal/web/static/index.html")
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("internal/web/static"))))

	// Serve static files (CSS, JS, etc.) at /static/
	fs := http.FileServer(http.Dir("internal/web/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Optional: API endpoint for future monitoring data
	r.HandleFunc("/api/audit", audit.AuditHandler).Methods("GET")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
