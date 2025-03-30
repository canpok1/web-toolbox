package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo, staticDir string) {
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Printf("static directory not found: %s", staticDir)
		return
	}

	// Serve static files from the "frontend" directory
	e.Static("/", staticDir)

	// Handle SPA routing
	e.GET("/planning-poker/*", func(c echo.Context) error {
		indexFilePath := filepath.Join(staticDir, "planning-poker", "index.html")
		if _, err := os.Stat(indexFilePath); os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "index.html not found")
		}
		return c.File(indexFilePath)
	})
}
