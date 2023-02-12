package db

import (
	"github.com/go-redis/redis/v8"
	"user-service/config"
)

type Database struct {
	redisCache *redis.Client
}

func Init() {
	cfg := config.GetConfig()
	initSqlDB(cfg)
	initRedis(cfg)
}

func initSqlDB(cfg *config.Configuration) {
	// Run Migrations based on models
}

func initRedis(cfg *config.Configuration) {
	// initialize redis cache
}
