package usecase

import (
	"context"

	"github.com/vincen320/user-service-graphql/model"
)

type HobbyUseCase interface {
	FindHobbies(ctx context.Context) (response []model.Hobby, err error)
}
