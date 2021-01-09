package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"calendar-backend/graph/model"
	"gorm.io/gorm"
)
// Resolver GraphQLのResolver定義
type Resolver struct{
	DB *gorm.DB
	schedules []*model.Schedule
	users []*model.User
}
