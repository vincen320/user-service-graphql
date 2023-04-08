package repository

import (
	"context"

	"github.com/vincen320/user-service-graphql/model"
)

type HobbyRepository interface {
	FindHobbies(ctx context.Context) (response []model.Hobby, err error)
}
