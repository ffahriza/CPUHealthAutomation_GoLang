package audit

import (
	"log"
)

type HealthReport struct {
	ServiceName string
	Status      string
	Detail      string
}

// Ini fungsi yang dimaksud:
func ScanServices(cfg interface{}) []HealthReport {
	log.Println("Scanning services...")
	return []HealthReport{
		{ServiceName: "database", Status: "healthy", Detail: "All good"},
		{ServiceName: "api", Status: "unhealthy", Detail: "Timeouts"},
	}
}
