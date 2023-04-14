package usecase

import (
	"context"

	"github.com/vincen320/user-service-graphql/model"
)

type (
	UserUseCase interface {
		FindUsers(ctx context.Context) (response []model.User, err error)
		FindUserByID(ctx context.Context, userID int64) (response []model.User, err error)
		CreateUser(ctx context.Context, request model.User) (user model.User, err error)
		Login(ctx context.Context, request model.UserLogin) (token string, err error)
	}
)
