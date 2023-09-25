package main

import (
	"fmt"
	"restEcho1/configs"
	"restEcho1/controller"
	"restEcho1/model"
	"restEcho1/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := model.InitModel(*config)
	model.Migrate(db)

	userModel := model.UsersModel{}
	userModel.Init(db)
	barangModel := model.BarangModel{}
	barangModel.Init(db)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	barangControll := controller.BarangController{}
	barangControll.InitUserController(barangModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userControll, *config)
	routes.RouteBarang(e, barangControll, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
