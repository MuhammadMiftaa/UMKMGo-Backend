package redis

import (
	"fmt"

	"sapaUMKM-backend/config/env"
	"sapaUMKM-backend/config/log"
	"sapaUMKM-backend/internal/utils/constant"

	"github.com/go-redis/redis/v8"
)

var RDB RedisRepository

type RedisRepository struct {
	Client *redis.Client
}

func SetupRedisDatabase(cfg env.Redis) {
	var db int
	if env.Cfg.Server.Mode == constant.DEVELOPMENT_MODE {
		db = 1
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RHost, cfg.RPort),
		DB:   db,
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}

	RDB.Client = rdb
}
