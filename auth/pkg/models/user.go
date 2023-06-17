package models

type User struct {
	ID       uint `gorm:"primaryKey;column:Id"`
	Login    string
	Email    string
	Password string
}
