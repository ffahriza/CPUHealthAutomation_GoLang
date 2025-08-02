package audit

import (
	"encoding/json"
	"net/http"

	"github.com/ffahriza/automation-smart-warehouse/internal/config"
)

var cfg config.Config // injected at init or via setter

// SetConfig inject config from main.go
func SetConfig(c config.Config) {
	cfg = c
}

func AuditHandler(w http.ResponseWriter, r *http.Request) {
	reports := ScanServices(cfg)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
