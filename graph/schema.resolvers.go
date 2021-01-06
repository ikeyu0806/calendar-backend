package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"calendar-backend/graph/generated"
	"calendar-backend/graph/model"
	"calendar-backend/infrastructure"
	"context"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *mutationResolver) CreateSchedule(ctx context.Context, input model.NewSchedule) (*model.Schedule, error) {
	schedule := &model.Schedule{
		Title:   input.Title,
		Content: input.Content,
		Memo:    input.Memo,
		StartAt: input.StartAt,
		EndAt:   input.EndAt,
	}
	db, err = infrastructure.GetDB()

	db.Create(&schedule)

	return schedule, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.UserToken, error) {
	secret := "safgvrebwabrq"

	input_password := input.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(*input_password), bcrypt.DefaultCost)

	password := string(hash)

	user := &model.User{
		Name:     input.Name,
		Password: &password,
		Email:    input.Email,
	}

	db, err = infrastructure.GetDB()

	if err = db.Create(&user).Error; err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": true,
		"name":  "ikegaya",
		"iat":   time.Now(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	userToken := &model.UserToken{
		ID:    user.ID,
		Name:  user.Name,
		Token: &tokenString,
	}

	return userToken, nil
}

func (r *queryResolver) Schedules(ctx context.Context) ([]*model.Schedule, error) {
	var schedules []*model.Schedule
	db, err = infrastructure.GetDB()
	db.Find(&schedules)
	return schedules, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	db  *gorm.DB
	err error
)
