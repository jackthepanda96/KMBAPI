package handler

import (
	"net/http"
	"restEcho1/features/users"
	"restEcho1/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.UserServiceInterface
}

func NewHandler(service users.UserServiceInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: service,
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil))
		}

		var serviceInput = new(users.User)
		serviceInput.Nama = input.Nama
		serviceInput.HP = input.HP
		serviceInput.Password = input.Password

		result, err := uh.s.Register(*serviceInput)

		if err != nil {
			c.Logger().Error("handler: input process error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil))
		}

		var response = new(RegisterResponse)
		response.Nama = result.Nama
		response.HP = result.HP

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", response))
	}
}
