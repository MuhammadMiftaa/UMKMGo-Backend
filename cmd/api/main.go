package main

import (
	"UMKMGo-backend/config/db"
	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/log"
	"UMKMGo-backend/config/redis"
	"UMKMGo-backend/config/storage"
	"UMKMGo-backend/interface/http/router"
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

	log.Info("Setup Redis Connection Start")
	redis.SetupRedisDatabase(env.Cfg.Redis) // Initialize the Redis connection
	log.Info("Setup Redis Connection Success")

	log.Info("Setup MinIO Connection Start")
	storage.SetupMinio(env.Cfg.Minio) // Initialize the MinIO connection
	log.Info("Setup MinIO Connection Success")

	log.Info("Starting UMKMGo API...")
}

func main() {
	defer log.Info("UMKMGo API stopped")

	r := router.SetupRouter() // Set up the HTTP router

	r.Listen(":" + env.Cfg.Server.Port)
	log.Info("Starting HTTP server on port " + env.Cfg.Server.Port)
}
