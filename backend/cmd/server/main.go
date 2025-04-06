package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/api/planningpoker"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/canpok1/web-toolbox/backend/internal/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Define environment variables with default values
	staticDir := getEnv("STATIC_DIR", "")
	redisAddress := getEnv("REDIS_ADDRESS", "redis:6379")
	port := getEnv("PORT", "8080")

	if staticDir == "" {
		log.Fatalf("Static directory path is empty, please set the STATIC_DIR environment variable")
	}
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory not found: path=%s", staticDir)
	}
	log.Printf("Static directory found: path=%s", staticDir)

	redisClient, err := redis.NewClient(redisAddress, "", 0, 24*time.Hour)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: address=%s, error=%v", redisAddress, err)
	}
	defer redisClient.Close()
	log.Printf("Success to connect to Redis: address=%s", redisAddress)

	webSocketHub := planningpoker.NewWebSocketHub()

	server := api.NewServer(redisClient, webSocketHub)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.RegisterHandlers(e, server)
	web.RegisterHandlers(e, staticDir)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("listen : %s\n", addr)

	log.Fatal(e.Start(addr))
}

// getEnv retrieves the value of an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
