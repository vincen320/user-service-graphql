package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/vincen320/user-service-graphql/usecase"
)

type hobbyHandler struct {
	graphqlSchema graphql.Schema
}

func NewHobbyHandler(hobbyUseCase usecase.HobbyUseCase, graphqlSchema graphql.Schema) *hobbyHandler {
	return &hobbyHandler{
		graphqlSchema: graphqlSchema,
	}
}
