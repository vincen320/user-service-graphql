package usecase

import (
	"context"

	"github.com/vincen320/user-service-graphql/model"
	"github.com/vincen320/user-service-graphql/repository"
	"github.com/vincen320/user-service-graphql/validator"
)

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) FindUsers(ctx context.Context) (response []model.User, err error) {
	response, err = u.userRepository.FindUsers(ctx)
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, len(response))
	for i, user := range response {
		userIDs[i] = user.ID
	}
	mapUserHobbies, err := u.userRepository.FindUserHobbies(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	for i, user := range response {
		if userHobbies, ok := mapUserHobbies[user.ID]; ok {
			user.Hobbies = userHobbies
		}
		response[i] = user
	}
	return
}

func (u *userUseCase) FindUserByID(ctx context.Context, userID int64) (response []model.User, err error) {
	return u.userRepository.FindUserByID(ctx, userID)
}

func (u *userUseCase) CreateUser(ctx context.Context, request model.User) (user model.User, err error) {
	err = validator.ValidateCreateUser(request)
	if err != nil {
		return
	}
	return u.userRepository.CreateUser(ctx, request)
}
