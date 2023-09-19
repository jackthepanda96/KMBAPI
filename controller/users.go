package controller

import (
	"net/http"
	"restEcho1/helper"
	"restEcho1/model"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model model.UsersModel
}

func (uc *UserController) InitUserController(um model.UsersModel) {
	uc.model = um
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
