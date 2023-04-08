package usecase

import (
	"context"

	"github.com/vincen320/user-service-graphql/model"
	"github.com/vincen320/user-service-graphql/repository"
)

type hobbyUseCase struct {
	hobbyRepositoryRepository repository.HobbyRepository
}

func NewHobbyUseCase(hobbyRepository repository.HobbyRepository) HobbyUseCase {
	return &hobbyUseCase{
		hobbyRepositoryRepository: hobbyRepository,
	}
}

func (h *hobbyUseCase) FindHobbies(ctx context.Context) (response []model.Hobby, err error) {
	return h.hobbyRepositoryRepository.FindHobbies(ctx)
}
