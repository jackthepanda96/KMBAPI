package service

import (
	"restEcho1/features/users"
	"restEcho1/helper"
)

type UserService struct {
	d users.UserDataInterface
	g helper.Generator
}

func New(data users.UserDataInterface, generator helper.Generator) users.UserServiceInterface {
	return &UserService{
		d: data,
		g: generator,
	}
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	newData.ID = us.g.GenerateUUID()
	result, err := us.d.Insert(newData)

	if err != nil {
		return nil, err
	}

	return result, nil
}
