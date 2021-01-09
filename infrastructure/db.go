package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
  db  *gorm.DB
  err error
)

// GetDB GormでMysqlに接続するメソッド
func GetDB() (*gorm.DB, error) {
	dsn := "root:pass@tcp(mysql:3306)/calendar?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db, err
}
