package routes

import (
	"restEcho1/configs"
	"restEcho1/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, uc controller.UserController, cfg configs.ProgramConfig) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	e.GET("/users", uc.MyProfile(), echojwt.JWT([]byte(cfg.Secret)))
}

func RouteBarang(e *echo.Echo, bc controller.BarangController) {
	e.POST("/barangs", bc.Insert())
	e.GET("/barangs", bc.GetBarangs(), middleware.JWT("yourKey"))
	e.DELETE("/barangs", bc.Delete())
	e.PUT("/barangs/:id", bc.Update())
}
