package controller

import (
	"restEcho1/configs"
	"restEcho1/model"
	"testing"
)

func TestRegister(t *testing.T) {
	// Setup Controller

	var cfg = configs.InitConfig()
	var gorm = model.InitModel(*cfg)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *cfg)

	// Setup ECHO
	// var e = echo.New()

	// var req = httptest.NewRequest(http.MethodPost, "/users", nil)
	// var res = httptest.NewRecorder()

	// var c = e.NewContext(req, res)

}
