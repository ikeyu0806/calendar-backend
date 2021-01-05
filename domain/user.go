package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int
	Name string
	Password string
	Email string
}

type Users []User
