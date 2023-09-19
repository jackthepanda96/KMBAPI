package main

import (
	"fmt"
	"restEcho1/configs"
	"restEcho1/controller"
	"restEcho1/model"
	"restEcho1/routes"

	"github.com/labstack/echo/v4"
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
	userControll.InitUserController(userModel)

	barangControll := controller.BarangController{}
	barangControll.InitUserController(barangModel)

	routes.RouteUser(e, userControll)
	routes.RouteBarang(e, barangControll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
