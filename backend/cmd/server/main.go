package main

import (
	"log"
	"net/http"

	"github.com/canpok1/web-toolbox/backend/internal/api"
	"github.com/canpok1/web-toolbox/backend/internal/redis"
)

func main() {
	redisClient, err := redis.NewClient("redis:6379", "", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	server := api.NewServer(redisClient)
	r := http.NewServeMux()
	h := api.HandlerWithOptions(server, api.StdHTTPServerOptions{
		BaseRouter: r,
	})

	addr := "0.0.0.0:8080"
	log.Printf("listen : %s\n", addr)

	s := &http.Server{
		Handler: h,
		Addr:    addr,
	}

	log.Fatal(s.ListenAndServe())
}
