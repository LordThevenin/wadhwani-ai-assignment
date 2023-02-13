package db

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"user-service/config"
	"user-service/entities"
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
	}
	initRedis(cfg)
	automigrateSqlDb()
}

func GetSqlDB() *gorm.DB {
	return db.sqlDB
}

func initSqlDB(cfg *config.Configuration) error {
	// Run migrations based on models
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ")/" + cfg.DbName + "?charset=utf8mb4&parseTime=True"
	sqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
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
}

func automigrateSqlDb() {
	err := db.sqlDB.AutoMigrate(&entities.User{})
	if err != nil {
		// Log error to automigrate user entity
	}
}
