package router

import (
	"dinozarl2-panel-api/internal/handlers"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, dbClient *dynamodb.Client) {
	wheelHandler := &handlers.WheelHandler{DBClient: dbClient}

	e.GET("/meow", handlers.Meow)
	e.GET("/meow/:count", handlers.Meow)
	e.POST("/meow", handlers.ConfigureMeow)
	e.GET("/wheels/:wheelId", wheelHandler.WheelGetById)
	e.DELETE("/wheels/:wheelId", wheelHandler.WheelsDeleteById)
	e.PUT("/wheels", wheelHandler.CreateOrUpdateWheel)
}
