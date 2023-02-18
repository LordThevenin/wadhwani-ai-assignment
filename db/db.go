package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"user-service/config"
	"user-service/entities"
	"user-service/utils"
)

type Database struct {
	redisCache *redis.Client
	sqlDB      *gorm.DB
}

var db *Database

func Init() {
	cfg := config.GetConfig()
	db = new(Database)
	err := initSqlDB(cfg)
	if err != nil {
		// Send error initialization failure
		utils.Logger().Errorf("Init (DB): failed to initialize sql db with error: %s", err.Error())
	}
	initRedis(cfg)
	automigrateSqlDb()
}

func GetSqlDB() *gorm.DB {
	return db.sqlDB
}

func GetRedisCache() *redis.Client {
	return db.redisCache
}

func initSqlDB(cfg *config.Configuration) error {
	// Run migrations based on models
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ")/" + cfg.DbName + "?charset=utf8mb4&parseTime=True"
	sqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
		utils.Logger().Errorf("initSqlDB: failed to initialize sql db with error: %s", err.Error())
		panic(err)
	}
	db.sqlDB, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDb}), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})

	var result int
	response := db.sqlDB.Raw("SELECT 1").Scan(&result)
	if response.Error != nil {
		return fmt.Errorf("DB fetch failed. Err: %w", response.Error)
	}
	return nil
}

func initRedis(cfg *config.Configuration) {
	// initialize redis cache
	ctx := context.Background()
	db.redisCache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "",          // no password set.
		DB:       cfg.RedisDB, // use default DB.
	})

	_, err := db.redisCache.Ping(ctx).Result()
	if err != nil {
		// Log error
		utils.Logger().Errorf("initRedis: failed to initialize redis cache with error: %s", err.Error())
	} else {
		// Log response
		utils.Logger().Debugf("initRedis: pong")
	}
}

func automigrateSqlDb() {
	err := db.sqlDB.AutoMigrate(&entities.User{})
	if err != nil {
		// Log error to automigrate user entity
		utils.Logger().Errorf("automigrateSqlDb: failed to automigrate entities.User with error: %s", err.Error())
	}
	err = db.sqlDB.AutoMigrate(&entities.AuthUser{})
	if err != nil {
		// Log error to automigrate authUser entity
		utils.Logger().Errorf("automigrateSqlDb: failed to automigrate entities.AuthUser with error: %s", err.Error())
	}
}
