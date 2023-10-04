package data

import (
	"restEcho1/features/users"

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
