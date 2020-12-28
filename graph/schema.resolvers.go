package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"gragql_demo/graph/generated"
	"gragql_demo/graph/model"
)

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		User:        users[0],
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A description",
		User:        users[1],
	},
}
var users = []*model.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Jon",
		Email:    "jon@gmail.com",
	},
}

// Meetups 类型中的User字段自定义展示函数，返回接口都在generated中生成好了
func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	user := new(model.User)

	for _, u := range users {
		if u.ID == obj.ID {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("user with id not exist")
	}

	return user, nil
}

// 更新数据的函数
func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

// 用于query的meetups字段的展示实现
func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return meetups, nil
}

// 用于每个user的meetups字段的自定义展示
func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var m []*model.Meetup

	for _, meetup := range meetups {
		if meetup.ID == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
