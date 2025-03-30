package main

import (
	"log"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/api/planningpoker"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
	"github.com/labstack/echo/v4"
)

func main() {
	redisAddress := "redis:6379"
	redisClient, err := redis.NewClient(redisAddress, "", 0)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %s, error: %v", redisAddress, err)
	}
	defer redisClient.Close()
	webSocketHub := planningpoker.NewWebSocketHub()

	server := api.NewServer(redisClient, webSocketHub)
	e := echo.New()
	api.RegisterHandlers(e, server)

	addr := "0.0.0.0:8080"
	log.Printf("listen : %s\n", addr)

	log.Fatal(e.Start(addr))
}
