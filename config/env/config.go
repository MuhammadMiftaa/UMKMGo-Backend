package env

import (
	"fmt"
	"os"
	"strconv"

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

	Redis struct {
		RHost string `env:"REDIS_HOST"`
		RPort string `env:"REDIS_PORT"`
	}

	Minio struct {
		Host        string `env:"MINIO_HOST"`
		AccessKey   string `env:"MINIO_ROOT_USER"`
		SecretKey   string `env:"MINIO_ROOT_PASSWORD"`
		MaxOpenConn int    `env:"MINIO_MAX_OPEN_CONN"`
		UseSSL      int    `env:"MINIO_USE_SSL"`
	}

	ZSMTP struct {
		ZSHost     string `env:"ZOHO_SMTP_HOST"`
		ZSPort     string `env:"ZOHO_SMTP_PORT"`
		ZSUser     string `env:"ZOHO_SMTP_USER"`
		ZSPassword string `env:"ZOHO_SMTP_PASSWORD"`
		ZSSecure   string `env:"ZOHO_SMTP_SECURE"`
		ZSAuth     bool   `env:"ZOHO_SMTP_AUTH"`
	}

	Config struct {
		Server   Server
		Database Database
		Redis    Redis
		Minio    Minio
		ZSMTP    ZSMTP
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

	// ! Load Redis configuration _____________________________
	if Cfg.Redis.RHost, ok = os.LookupEnv("REDIS_HOST"); !ok {
		missing = append(missing, "REDIS_HOST env is not set")
	}
	if Cfg.Redis.RPort, ok = os.LookupEnv("REDIS_PORT"); !ok {
		missing = append(missing, "REDIS_PORT env is not set")
	}
	// ! ______________________________________________________

	// ! Load MinIO configuration _____________________________
	if Cfg.Minio.Host, ok = os.LookupEnv("MINIO_HOST"); !ok {
		missing = append(missing, "MINIO_HOST env is not set")
	}
	if Cfg.Minio.AccessKey, ok = os.LookupEnv("MINIO_ROOT_USER"); !ok {
		missing = append(missing, "MINIO_ROOT_USER env is not set")
	}
	if Cfg.Minio.SecretKey, ok = os.LookupEnv("MINIO_ROOT_PASSWORD"); !ok {
		missing = append(missing, "MINIO_ROOT_PASSWORD env is not set")
	}
	if val, ok := os.LookupEnv("MINIO_MAX_OPEN_CONN"); !ok {
		missing = append(missing, "MINIO_MAX_OPEN_CONN env is not set")
	} else {
		var err error
		if Cfg.Minio.MaxOpenConn, err = strconv.Atoi(val); err != nil {
			missing = append(missing, fmt.Sprintf("MINIO_MAX_OPEN_CONN must be int, got %s", val))
		}
	}
	if val, ok := os.LookupEnv("MINIO_USE_SSL"); !ok {
		missing = append(missing, "MINIO_USE_SSL env is not set")
	} else {
		var err error
		if Cfg.Minio.UseSSL, err = strconv.Atoi(val); err != nil {
			missing = append(missing, fmt.Sprintf("MINIO_USE_SSL must be int, got %s", val))
		}
	}
	// ! ______________________________________________________

	// ! Load Zoho SMTP configuration __________________________
	if Cfg.ZSMTP.ZSHost, ok = os.LookupEnv("ZOHO_SMTP_HOST"); !ok {
		missing = append(missing, "ZOHO_SMTP_HOST env is not set")
	}
	if Cfg.ZSMTP.ZSPort, ok = os.LookupEnv("ZOHO_SMTP_PORT"); !ok {
		missing = append(missing, "ZOHO_SMTP_PORT env is not set")
	}
	if Cfg.ZSMTP.ZSUser, ok = os.LookupEnv("ZOHO_SMTP_USER"); !ok {
		missing = append(missing, "ZOHO_SMTP_USER env is not set")
	}
	if Cfg.ZSMTP.ZSPassword, ok = os.LookupEnv("ZOHO_SMTP_PASSWORD"); !ok {
		missing = append(missing, "ZOHO_SMTP_PASSWORD env is not set")
	}
	if Cfg.ZSMTP.ZSSecure, ok = os.LookupEnv("ZOHO_SMTP_SECURE"); !ok {
		missing = append(missing, "ZOHO_SMTP_SECURE env is not set")
	}
	if zohoAuth, ok := os.LookupEnv("ZOHO_SMTP_AUTH"); !ok {
		missing = append(missing, "ZOHO_SMTP_AUTH env is not set")
	} else {
		Cfg.ZSMTP.ZSAuth = zohoAuth == "true"
	}
	// ! ______________________________________________________

	return missing, nil
}
