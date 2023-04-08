package presenter

import (
	"github.com/graphql-go/graphql"
	"github.com/vincen320/user-service-graphql/usecase"
)

type hobbyPresenter struct {
	graphqlSchema graphql.Schema
}

func NewHobbyPresenter(hobbyUseCase usecase.HobbyUseCase, graphqlSchema graphql.Schema) *hobbyPresenter {
	return &hobbyPresenter{
		graphqlSchema: graphqlSchema,
	}
}
