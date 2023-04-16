package repository

import (
	"context"

	"github.com/vincen320/user-service-graphql/model"
)

type UserRepository interface {
	FindUsers(ctx context.Context) (response []model.User, err error)
	FindUserByID(ctx context.Context, userID int64) (response []model.User, err error)
	FindUserByEmail(ctx context.Context, userEmail string) (response []model.User, err error)
	CreateUser(ctx context.Context, request model.User) (user model.User, err error)
	FindUserHobbies(ctx context.Context, userIDs []int64) (response map[int64][]model.Hobby, err error)
}
