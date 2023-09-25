package controller

import (
	"net/http"
	"restEcho1/configs"
	"restEcho1/helper"
	"restEcho1/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	cfg   configs.ProgramConfig
	model model.UsersModel
}

func (uc *UserController) InitUserController(um model.UsersModel, c configs.ProgramConfig) {
	uc.model = um
	uc.cfg = c
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.Users{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		var res = uc.model.Register(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res))
	}
}

func (uc *UserController) RefreshToken() echo.HandlerFunc {
	return func(c echo.Context) error {
		type RefreshInput struct {
			Token string `json:"access_token" form:"access_token"`
		}
		var input = RefreshInput{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		var currentToken = c.Get("user").(*jwt.Token)

		var response = helper.RefereshJWT(&jwt.Token{
			Raw: input.Token,
		}, currentToken)
		return c.JSON(http.StatusOK, helper.FormatResponse("success", response))
	}
}

func (uc *UserController) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var token = c.Get("user")
		var jwtClaims = helper.ExtractToken(token.(*jwt.Token))
		var mapClaims = jwtClaims.(jwt.Claims).(jwt.MapClaims)
		logrus.Info(mapClaims["id"])

		return c.JSON(http.StatusOK, "ok")
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = model.LoginModel{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("invalid user input", nil))
		}

		var res = uc.model.Login(input.HP, input.Sandi)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data, something happend", nil))
		}

		if res.Id == "" {
			return c.JSON(http.StatusNotFound, helper.FormatResponse("data not found", nil))
		}

		var jwtToken = helper.GenerateJWT(uc.cfg.Secret, uc.cfg.RefreshSecret, res.Id)

		if jwtToken == nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("cannot process data", nil))
		}

		jwtToken["info"] = res

		return c.JSON(http.StatusOK, helper.FormatResponse("success", jwtToken))
	}
}
