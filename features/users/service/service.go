package service

import (
	"errors"
	"restEcho1/features/users"
	"restEcho1/helper"
)

type UserService struct {
	d users.UserDataInterface
	g helper.GeneratorInterface
}

func New(data users.UserDataInterface, generator helper.GeneratorInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		g: generator,
	}
}

func (us *UserService) Register(newData users.User) (users.User, error) {
	newID, err := us.g.GenerateUUID()
	if err != nil {
		return users.User{}, errors.New("id generator failed")
	}

	newData.ID = newID
	result, err := us.d.Insert(newData)
	if err != nil {
		return users.User{}, errors.New("insert process failed")
	}

	return result, nil
}
