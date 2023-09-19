package controller

import (
	"fmt"
	"net/http"
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
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "invalid user input",
			})
		}
		fmt.Println(input)
		var res = uc.model.Register(input)

		if res == nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "cannot process data, something happend",
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    res,
		})
	}
}
