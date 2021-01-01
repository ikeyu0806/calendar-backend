package domain

import (
	"time"
	"gorm.io/gorm"
)

type Schedule struct {
  gorm.Model
	UserId int
	Title string
	Content string
	Memo string `gorm:"type:text"`
	StartAt time.Time
	EndAt time.Time
}

type Schedules []Schedule
