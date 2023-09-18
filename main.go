package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	id    string
	Nama  string `json:"name" form:"name"`
	HP    string `json:"hp" form:"hp"`
	Sandi string `json:"password" form:"password"`
}

func (u *Users) SetPassword() {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Sandi), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
	}
	u.Sandi = string(encryptedPassword)
}

func (u *Users) GenerateID() {
	if u.id == "" {
		u.id = uuid.NewString()
		return
	}
	fmt.Println("ID sudah ada")
}

var (
	ListUser []Users
)

// func Register(newUser Users) {
// 	newUser.SetPassword()
// 	newUser.GenerateID()
// 	ListUser = append(ListUser, newUser)
// }

func SetResponse(message string, data any) map[string]any {
	if data == nil {
		return map[string]any{
			"message": message,
		}
	}

	return map[string]any{
		"message": message,
		"data":    data,
	}
}

func Register(c echo.Context) error {
	var input = Users{}
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(fmt.Sprint("error ketika parsing data -", err.Error()), nil))
	}
	input.GenerateID()
	ListUser = append(ListUser, input)
	return c.JSON(http.StatusOK, SetResponse("sukses", nil))
}

func Register2() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = Users{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, SetResponse(fmt.Sprint("error ketika parsing data -", err.Error()), nil))
		}
		input.GenerateID()
		ListUser = append(ListUser, input)
		return c.JSON(http.StatusOK, SetResponse("sukses", nil))
	}
}

func GetDataByIndex(c echo.Context) error {
	var pathParam1 = c.Param("idx")
	idx, err := strconv.Atoi(pathParam1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
	}

	if idx < 0 || idx > len(ListUser)-1 {
		return c.JSON(http.StatusNotFound, SetResponse("index tidak terdaftar", nil))
	}

	var res = ListUser[idx]

	return c.JSON(http.StatusOK, SetResponse("sukses", res))
}

func Update(c echo.Context) error {
	var pathParam1 = c.Param("idx")
	var updateData = Users{}
	idx, err := strconv.Atoi(pathParam1)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
	}

	err = c.Bind(&updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, SetResponse(fmt.Sprint("error ketika parsing data -", err.Error()), nil))
	}

	ListUser[idx].HP = updateData.HP
	ListUser[idx].Nama = updateData.Nama

	return c.JSON(http.StatusOK, SetResponse("sukses", nil))

}

func Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var pathParam1 = c.Param("idx")
		idx, err := strconv.Atoi(pathParam1)

		if err != nil {
			return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
		}

		if idx < 0 || idx > len(ListUser)-1 {
			return c.JSON(http.StatusNotFound, SetResponse("index tidak terdaftar", nil))
		}

		ListUser[idx] = Users{}
		return c.JSON(http.StatusNoContent, nil)
	}
}

func main() {
	e := echo.New()

	// e.GET("/users", func(c echo.Context) error {
	// 	if len(ListUser) == 0 {
	// 		return c.JSON(http.StatusNotFound,
	// 			SetResponse("data tidak ditemukan", nil))

	// 	}
	// 	return c.JSON(200, SetResponse("sukses", ListUser))

	// })
	// e.POST("/users", Register)
	// e.PUT("/users/:idx", Update)

	var users = e.Group("/users")
	users.GET("", func(c echo.Context) error {

		if len(ListUser) == 0 {
			return c.JSON(http.StatusNotFound,
				SetResponse("data tidak ditemukan", nil))

		}
		return c.JSON(200, SetResponse("sukses", ListUser))

	})
	users.POST("", Register)
	users.PUT("/:idx", Update)
	users.GET("/:idx", GetDataByIndex)
	users.DELETE("/:idx", Delete())

	e.Logger.Fatal(e.Start(":8000").Error())
}
