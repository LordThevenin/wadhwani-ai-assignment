package main

import (
	"user-service/config"
	"user-service/db"
	"user-service/server"
)

func main() {
	// Initialize Config
	config.Init()
	// Initialize Components
	db.Init()
	// Start server
	server.Start()
}
