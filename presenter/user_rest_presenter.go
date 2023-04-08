package presenter

import (
	"github.com/graphql-go/graphql"
	"github.com/vincen320/user-service-graphql/usecase"
)

type userPresenter struct {
	userUseCase usecase.UserUseCase
}

func NewUserPresenter(userUseCase usecase.UserUseCase, graphqlSchema graphql.Schema) *userPresenter {
	return &userPresenter{
		userUseCase: userUseCase,
	}
}
