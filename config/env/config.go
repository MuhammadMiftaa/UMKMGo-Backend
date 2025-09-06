package env

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Server struct {
		Mode         string `env:"MODE"`
		Port         string `env:"PORT"`
		JWTSecretKey string `env:"JWT_SECRET_KEY"`
	}

	Database struct {
		DBHost     string `env:"DB_HOST"`
		DBPort     string `env:"DB_PORT"`
		DBUser     string `env:"DB_USER"`
		DBPassword string `env:"DB_PASSWORD"`
		DBName     string `env:"DB_NAME"`
	}

	Config struct {
		Server   Server
		Database Database
	}
)

var Cfg Config

func LoadNative() ([]string, error) {
	var ok bool
	var missing []string

	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	// ! Load Server configuration ____________________________
	if Cfg.Server.Mode, ok = os.LookupEnv("MODE"); !ok {
		missing = append(missing, "MODE env is not set")
	}
	if Cfg.Server.Port, ok = os.LookupEnv("PORT"); !ok {
		missing = append(missing, "PORT env is not set")
	}
	if Cfg.Server.JWTSecretKey, ok = os.LookupEnv("JWT_SECRET_KEY"); !ok {
		missing = append(missing, "JWT_SECRET_KEY env is not set")
	}
	// ! ______________________________________________________

	// ! Load Database configuration __________________________
	if Cfg.Database.DBUser, ok = os.LookupEnv("DB_USER"); !ok {
		missing = append(missing, "DB_USER env is not set")
	}
	if Cfg.Database.DBHost, ok = os.LookupEnv("DB_HOST"); !ok {
		missing = append(missing, "DB_HOST env is not set")
	}
	if Cfg.Database.DBPort, ok = os.LookupEnv("DB_PORT"); !ok {
		missing = append(missing, "DB_PORT env is not set")
	}
	if Cfg.Database.DBName, ok = os.LookupEnv("DB_NAME"); !ok {
		missing = append(missing, "DB_NAME env is not set")
	}
	if Cfg.Database.DBPassword, ok = os.LookupEnv("DB_PASSWORD"); !ok {
		missing = append(missing, "DB_PASSWORD env is not set")
	}
	// ! ______________________________________________________

	return missing, nil
}
