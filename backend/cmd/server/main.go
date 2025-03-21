package main

import (
	"log"
	"net/http"

	"github.com/canpok1/web-toolbox/backend/internal/api"
)

func main() {
	server := api.NewServer()
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
