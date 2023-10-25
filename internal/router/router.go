package router

import (
	"dinozarl2-panel-api/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/meow", handlers.Meow)
	e.GET("/meow/:count", handlers.Meow)
	e.GET("/wheels/:wheelId", handlers.WheelGetById)
	e.DELETE("/wheels/:wheelId", handlers.WheelsDeleteById)
}
