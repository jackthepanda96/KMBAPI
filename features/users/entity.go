package users

import "github.com/labstack/echo/v4"

type User struct {
	ID       string
	Nama     string
	HP       string
	Password string
}

type UserCredential struct {
	Nama   string
	Access map[string]any
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}
type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(hp string, password string) (*UserCredential, error)
}
type UserDataInterface interface {
	Insert(newData User) (*User, error)
	Login(hp string, password string) (*User, error)
}
