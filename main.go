package main

import (
	"fmt"
	"net/http"
	"restEcho1/configs"
	"restEcho1/controller"
	"restEcho1/model"
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
	var config = configs.InitConfig()
	fmt.Println(config)

	db := model.InitModel(*config)
	model.Migrate(db)

	userModel := model.UsersModel{}
	userModel.Init(db)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel)

	var users = e.Group("/users")
	users.POST("", userControll.Register())
	// users.PUT("/:idx", Update)
	// users.GET("/:idx", GetDataByIndex)
	// users.DELETE("/:idx", Delete())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
