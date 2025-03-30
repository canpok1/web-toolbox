package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, publicDir string) {
	if _, err := os.Stat(publicDir); os.IsNotExist(err) {
		log.Printf("public directory not found: %s", publicDir)
		return
	}

	log.Printf("public directory found: %s", publicDir)

	// Serve static files from the "frontend" directory
	e.Static("/", publicDir)

	// Handle SPA routing
	e.GET("/planning-poker/*", func(c echo.Context) error {
		indexFilePath := filepath.Join(publicDir, "planning-poker", "index.html")
		if _, err := os.Stat(indexFilePath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "index.html not found")
		}
		return c.File(indexFilePath)
	})
}
