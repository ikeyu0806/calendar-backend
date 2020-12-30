package main

import (
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Schedule struct {
  gorm.Model
	UserId int `json:"movie_id"`
	StartAt time.Time
	EndAt time.Time
}

var (
  db  *gorm.DB
  err error
)

func main() {
	dsn := "root:pass@tcp(localhost:3306)/calendar?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	s := Schedule{UserId: 1, StartAt: time.Now(), EndAt: time.Now()}

	if err != nil {
		db.Create(&s)
	}
}
