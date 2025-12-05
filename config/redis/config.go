package redis

import (
	"fmt"

	"UMKMGo-backend/config/env"
	"UMKMGo-backend/config/log"
	"UMKMGo-backend/internal/utils/constant"

	redisPackage "github.com/go-redis/redis/v8"
)

var redis redisInstance

type redisInstance struct {
	Client *redisPackage.Client
}

func SetupRedisDatabase(cfg env.Redis) {
	var db int
	if env.Cfg.Server.Mode == constant.DEVELOPMENT_MODE {
		db = 1
	}

	rdb := redisPackage.NewClient(&redisPackage.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RHost, cfg.RPort),
		DB:   db,
	})

	_, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}

	redis.Client = rdb
}

func GetRedisRepository() RedisRepository {
	if redis.Client == nil {
		log.Log.Fatal("Redis client is not initialized. Please call SetupRedisDatabase first.")
	}
	return &redis
}
