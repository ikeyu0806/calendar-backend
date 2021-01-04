package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"calendar-backend/graph/model"
	"gorm.io/gorm"
)
type Resolver struct{
	DB *gorm.DB
	todos []*model.Todo
	schedules []*model.Schedule
	users []*model.User
}
