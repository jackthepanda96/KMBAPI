package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Users struct {
	Id      string `gorm:"primaryKey;type:varchar(20)"`
	Nama    string `gorm:"type:varchar(255)"`
	HP      string `gorm:"type:varchar(13);uniqueIndex"`
	Sandi   string
	Barangs []Barang `gorm:"foreignKey:Pemilik;references:Id"`
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

func (um *UsersModel) GetAll() []Users {
	var listUser = []Users{}
	if err := um.db.Find(&listUser).Error; err != nil {
		logrus.Error("Model : Insert data error, ", err.Error())
		return nil
	}

	return listUser
}
