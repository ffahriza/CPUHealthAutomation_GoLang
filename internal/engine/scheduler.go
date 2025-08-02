package engine

import (
	"log"
	"time"

	"github.com/ffahriza/automation-smart-warehouse/internal/audit"
	"github.com/ffahriza/automation-smart-warehouse/internal/config"
	"github.com/ffahriza/automation-smart-warehouse/internal/service"
)

func StartScheduler(cfg config.Config) {
	log.Println("Starting Audit & Healing Scheduler")

	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		healthReports := audit.ScanServices(cfg)

		for _, report := range healthReports {
			if report.Status == "unhealthy" {
				log.Printf("Service %s is unhealthy. Triggering remediation...\n", report.ServiceName)
				service.TriggerHealing(report.ServiceName)
			}
		}
	}
}
