package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"calendar-backend/graph/generated"
	"calendar-backend/graph/model"
	"calendar-backend/infrastructure"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

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

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}
	db, err = infrastructure.GetDB()

	db.Create(&user)

	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
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
