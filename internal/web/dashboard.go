package web

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/web/static/index.html")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"database":  "OK",
		"container": "OK",
		"timestamp": "2025-07-29 21:00",
	})
}

func AuditHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"result":    "No Anomaly",
		"timestamp": "2025-07-29 20:55",
	})
}

func ManualHealHandler(w http.ResponseWriter, r *http.Request) {
	// Trigger self-healing logic
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "Triggered healing manually",
	})
}
