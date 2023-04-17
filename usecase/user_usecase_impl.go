package usecase

import (
	"context"
	"fmt"
	"net/http"

	cError "github.com/vincen320/user-service-graphql/helper/error"
	"github.com/vincen320/user-service-graphql/model"
	"github.com/vincen320/user-service-graphql/repository"
	"github.com/vincen320/user-service-graphql/validator"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	request.Password = string(hashedPassword)
	return u.userRepository.CreateUser(ctx, request)
}

func (u *userUseCase) Login(ctx context.Context, request model.UserLogin) (token string, err error) {
	err = validator.ValidateUserLogin(request)
	if err != nil {
		return
	}
	user, err := u.userRepository.FindUserByEmail(ctx, request.Email)
	if err != nil {
		return
	}
	if len(user) == 0 {
		err = cError.New(http.StatusNotFound, fmt.Sprintf("user with %s email not found", request.Email), "user not found")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(request.Password))
	if err != nil {
		err = cError.New(http.StatusUnauthorized, "wrong password", err.Error())
		return
	}
	return user[0].GenerateJWTToken()
}
