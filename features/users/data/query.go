package data

import (
	"restEcho1/features/users"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserData struct {
	gorm *gorm.DB
}

func New(g *gorm.DB) users.UserDataInterface {
	return &UserData{
		gorm: g,
	}
}

func (ud *UserData) Insert(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.ID = newData.ID
	dbData.HP = newData.HP
	dbData.Nama = newData.Nama
	dbData.Password = newData.Password

	if err := ud.gorm.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ud *UserData) Login(hp string, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.HP = hp
	dbData.Password = password

	if err := ud.gorm.Where("hp = ? AND password = ?", dbData.HP, dbData.Password).First(dbData).Error; err != nil {
		logrus.Info("db error:", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Nama = dbData.Nama
	result.HP = dbData.HP

	return result, nil
}
