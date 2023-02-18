package main

import (
	"user-service/config"
	"user-service/db"
	"user-service/server"
	"user-service/utils"
)

func main() {
	// Initialize logger
	utils.InitLogger()
	// Initialize Config
	config.Init()
	// Initialize Components
	db.Init()
	// Start server
	server.Start()
}
