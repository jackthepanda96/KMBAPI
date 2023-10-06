package service

import (
	"errors"
	"restEcho1/features/users"
	"restEcho1/helper"
	"strings"
)

type UserService struct {
	d users.UserDataInterface
	g helper.GeneratorInterface
	j helper.JWTInterface
}

func New(data users.UserDataInterface, generator helper.GeneratorInterface, jwt helper.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		g: generator,
		j: jwt,
	}
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	newID, err := us.g.GenerateUUID()
	if err != nil {
		return nil, errors.New("id generator failed")
	}

	newData.ID = newID
	result, err := us.d.Insert(newData)
	if err != nil {
		return nil, errors.New("insert process failed")
	}

	return result, nil
}

func (us *UserService) Login(hp string, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(hp, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("process failed")
	}

	tokenData := us.j.GenerateJWT(result.ID)

	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Nama = result.Nama
	response.Access = tokenData

	return response, nil
}
