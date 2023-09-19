package routes

import (
	"restEcho1/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserController) {
	e.POST("/users", uc.Register())
}

func RouteBarang(e *echo.Echo, bc controller.BarangController) {
	e.POST("/barangs", bc.Insert())
	e.GET("/barangs", bc.GetBarangs())
	e.DELETE("/barangs", bc.Delete())
}
