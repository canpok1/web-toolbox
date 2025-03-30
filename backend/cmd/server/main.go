package main

import (
	"flag"
	"log"
	"os"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/api/planningpoker"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/canpok1/web-toolbox/backend/internal/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Define command-line flags
	staticDir := flag.String("static-dir", "../frontend", "Path to the static files directory")
	flag.Parse()

	// Check if the static directory exists
	if _, err := os.Stat(*staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory not found: %s", *staticDir)
	}

	redisAddress := "redis:6379"
	redisClient, err := redis.NewClient(redisAddress, "", 0)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s, error: %v", redisAddress, err)
	}
	defer redisClient.Close()
	webSocketHub := planningpoker.NewWebSocketHub()

	server := api.NewServer(redisClient, webSocketHub)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api.RegisterHandlers(e, server)
	web.RegisterHandlers(e, *staticDir)

	addr := "0.0.0.0:8080"
	log.Printf("listen : %s\n", addr)

	log.Fatal(e.Start(addr))
}
