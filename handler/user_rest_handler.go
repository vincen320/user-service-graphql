package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/vincen320/user-service-graphql/usecase"
)

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase, graphqlSchema graphql.Schema) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}
