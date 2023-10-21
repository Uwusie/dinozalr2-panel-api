package server

import (
	"dinozarl2-panel-api/internal/router"
	"log"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	router.SetupRoutes(e)
	err := e.Start(":2137")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
