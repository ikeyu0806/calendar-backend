package main

import (
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Schedule struct {
  gorm.Model
	UserId int
	StartAt time.Time
	EndAt time.Time
}

var (
  db  *gorm.DB
  err error
)


func gormConnectTest() () {
	dsn := "root:pass@tcp(mysql:3306)/calendar?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	schedule := Schedule{UserId: 1, StartAt: time.Now(), EndAt: time.Now()}

	result := db.Create(&schedule)

	log.Println(result.Error)
}

func main() {
	gormConnectTest()
}
