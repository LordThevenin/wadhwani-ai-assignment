package server

import (
	"strconv"
	"user-service/config"
)

// Start initialized the server and begins to listen at the configured port
func Start() {
	cfg := config.GetConfig()
	r := NewRouter()
	r.Run(":" + strconv.Itoa(cfg.Port))
}
