package presenter

import (
	"github.com/graphql-go/graphql"
	"github.com/vincen320/user-service-graphql/model"
	"github.com/vincen320/user-service-graphql/usecase"
)

type hobbyGQL struct {
	hobbyUseCase usecase.HobbyUseCase
}

func NewHobbyGQL(hobbyUseCase usecase.HobbyUseCase) *hobbyGQL {
	return &hobbyGQL{
		hobbyUseCase: hobbyUseCase,
	}
}

func (h *hobbyGQL) GetHobbies() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(model.HobbyType),
		Resolve: func(p graphql.ResolveParams) (response interface{}, err error) {
			response, err = h.hobbyUseCase.FindHobbies(p.Context)
			return
		},
	}
}
