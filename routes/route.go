package routes

import (
	"sesi6/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) *echo.Echo {

	e.GET("/news", controllers.NewsController)
	e.GET("/news/:id", controllers.DetailNewsController)
	e.POST("/news", controllers.CreateNewsController)
	e.PUT("/news/:id", controllers.UpdateController)
	return e
}
