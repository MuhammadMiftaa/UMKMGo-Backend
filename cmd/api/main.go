package main

import (
	"sapaUMKM-backend/config/db"
	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/config/log"
	"sapaUMKM-backend/interface/http/router"
)

func init() {
	log.SetupLogger() // Initialize the logger configuration

	if missing, err := env.LoadNative(); err != nil {
		log.Log.Fatalf("Failed to load environment variables: %v", err)
	} else {
		if len(missing) > 0 {
			for _, envVar := range missing {
				log.Warn("Missing environment variable: " + envVar)
			}
		}
	}

	log.Info("Setup Database Connection Start")
	db.SetupDatabase(env.Cfg.Database) // Initialize the database connection and run migrations
	log.Info("Setup Database Connection Success")

	log.Info("Starting sapaUMKM API...")
}

func main() {
	defer log.Info("sapaUMKM API stopped")

	r := router.SetupRouter() // Set up the HTTP router
	r.Run(":" + env.Cfg.Server.Port)
	log.Info("Starting HTTP server on port " + env.Cfg.Server.Port)
}
