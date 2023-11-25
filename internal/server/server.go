package server

import (
	"context"
	"dinozarl2-panel-api/internal/database"
	"dinozarl2-panel-api/internal/rabbitmq"
	"dinozarl2-panel-api/internal/router"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	dynamoDBClient := database.NewDBClient()
	router.SetupRoutes(e, dynamoDBClient)

	go func() {
		if err := e.Start(":2137"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	rabbitmq.CloseChannel()
	rabbitmq.CloseConnection()
}
