package infrastructure

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// GetDB GormでMysqlに接続するメソッド
func GetDB() (*gorm.DB, error) {
	dsn := os.Getenv(("DATA_SOURCE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	} else {
		log.Print("sucess DB connect")
	}

	return db, err
}
