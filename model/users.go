package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Users struct {
	Id    string
	Nama  string `json:"name" form:"name"`
	HP    string `json:"hp" form:"hp"`
	Sandi string `json:"password" form:"password"`
}

func (u *Users) GenerateID() {
	if u.Id == "" {
		u.Id = uuid.NewString()
		return
	}
	fmt.Println("ID sudah ada")
}

type UsersModel struct {
	db *gorm.DB
}

func (um *UsersModel) Init(db *gorm.DB) {
	um.db = db
}

func (um *UsersModel) Register(newUser Users) *Users {
	newUser.GenerateID()
	if err := um.db.Create(&newUser).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return &newUser
}
