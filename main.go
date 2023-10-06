package main

import (
	"fmt"
	"restEcho1/configs"
	"restEcho1/features/users/data"
	"restEcho1/features/users/handler"
	"restEcho1/features/users/service"
	"restEcho1/helper"
	"restEcho1/routes"
	"restEcho1/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = configs.InitConfig()

	db := database.InitDB(*config)
	database.Migrate(db)

	userModel := data.New(db)
	generator := helper.NewGenerator()
	jwtInterface := helper.New(config.Secret, config.RefreshSecret)
	userServices := service.New(userModel, generator, jwtInterface)
	// barangModel := model.NewBarangModel(db)

	userControll := handler.NewHandler(userServices)

	// barangControll := controller.NewBarangControllInterface(barangModel)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userControll, *config)
	// routes.RouteBarang(e, barangControll, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
