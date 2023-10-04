package data

type User struct {
	ID       string `gorm:"varchar(255);primaryKey;"`
	Nama     string
	HP       string
	Password string
}
