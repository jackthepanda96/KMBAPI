package users

import "github.com/labstack/echo/v4"

type User struct {
	ID       string
	Nama     string
	HP       string
	Password string
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
}
type UserServiceInterface interface {
	Register(newData User) (*User, error)
}
type UserDataInterface interface {
	Insert(newData User) (*User, error)
}
