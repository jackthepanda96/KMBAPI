package routes

import (
	"restEcho1/configs"
	"restEcho1/features/users"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc users.UserHandlerInterface, cfg configs.ProgramConfig) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	// e.GET("/users", uc.MyProfile(), echojwt.JWT([]byte(cfg.Secret)))
	// // e.GET("/users/:id",)
	// e.POST("/refresh", uc.RefreshToken(), echojwt.JWT([]byte(cfg.RefreshSecret)))
}

// func RouteBarang(e *echo.Echo, bc controller.BarangControllInterface, cfg configs.ProgramConfig) {
// 	var barang = e.Group("/barangs")
// 	// barang.Use(echojwt.JWT([]byte(cfg.Secret)))

// 	barang.POST("", bc.Insert())
// 	barang.GET("", bc.GetBarangs())
// 	barang.DELETE("", bc.Delete())
// 	barang.PUT("/:id", bc.Update())
// }
