package handler

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/vincen320/user-service-graphql/helper"
	cError "github.com/vincen320/user-service-graphql/helper/error"
	"github.com/vincen320/user-service-graphql/model"
	"github.com/vincen320/user-service-graphql/usecase"
)

type (
	userGQL struct {
		userUseCase usecase.UserUseCase
	}
)

func NewUserGQL(userUseCase usecase.UserUseCase) *userGQL {
	return &userGQL{
		userUseCase: userUseCase,
	}
}

func (u *userGQL) GetUsers() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(model.UserType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (response interface{}, err error) {
			userID, ok := p.Args["id"].(int)
			if userID > 0 && ok {
				response, err = u.userUseCase.FindUserByID(p.Context, int64(userID))
				return
			}
			response, err = u.userUseCase.FindUsers(p.Context)
			return
		},
	}
}

func (u *userGQL) CreateUser() *graphql.Field {
	return &graphql.Field{
		Type: model.UserType,
		Args: graphql.FieldConfigArgument{
			"user": &graphql.ArgumentConfig{
				Type: model.CreateUserParam,
			},
		},
		Resolve: func(p graphql.ResolveParams) (response interface{}, err error) {
			request, err := helper.DecodeRequest[model.User](p.Args["user"])
			if err != nil {
				return response, cError.New(http.StatusBadRequest, "invalid payload", "invalid payload received")
			}
			response, err = u.userUseCase.CreateUser(p.Context, request)
			return
		},
	}
}

func (u *userGQL) Login() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "token",
			Description: "token response / this should be create separate from type for pretty, but this is for test",
			Fields: graphql.Fields{
				"token": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						tokenString, ok := p.Source.(string)
						if ok {
							return tokenString, nil
						}
						return "", nil
					},
				},
			},
		}),
		Args: graphql.FieldConfigArgument{
			"login": &graphql.ArgumentConfig{
				Type: model.UserLoginParam,
			},
		},
		Resolve: func(p graphql.ResolveParams) (response interface{}, err error) {
			request, err := helper.DecodeRequest[model.UserLogin](p.Args["login"])
			if err != nil {
				return response, cError.New(http.StatusBadRequest, "invalid payload", "invalid payload received")
			}
			response, err = u.userUseCase.Login(p.Context, request)
			return
		},
	}
}
