package database

import (
	"time"
	"gorm.io/gorm"
	"calendar-backend/infrastructure"
	"calendar-backend/domain"
)

var (
  db  *gorm.DB
  err error
)

func StoreTest() {
	db, err = infrastructure.GetDB()
	schedule := domain.Schedule{UserId: 1, StartAt: time.Now(), EndAt: time.Now()}

	db.Create(&schedule)
}
