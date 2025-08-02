package service

import "log"

func TriggerHealing(serviceName string) {
	switch serviceName {
	case "api":
		log.Println("Restarting API service...")
		// healing logika: misalnya restart container
	case "database":
		log.Println("Reconnecting to database...")
	default:
		log.Println("Unknown service for healing.")
	}
}
