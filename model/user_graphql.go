package model

import (
	"github.com/graphql-go/graphql"
)

var (
	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "user",
		Description: "object of user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"salary": &graphql.Field{
				Type: graphql.Float,
			},
			"hobbies": &graphql.Field{
				Type: graphql.NewList(HobbyType),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user, userOk := p.Source.(User)
					limit, limitOk := p.Args["first"].(int)
					if userOk && limitOk && len(user.Hobbies) > 0 {
						hobbiesCount := len(user.Hobbies)
						if limit > hobbiesCount {
							limit = hobbiesCount
						}
						user.Hobbies = user.Hobbies[:limit]
					}
					return user.Hobbies, nil
				},
			},
		},
	})

	CreateUserParam = graphql.NewInputObject(graphql.InputObjectConfig{
		Name:        "CreateUserParam",
		Description: "create user input params",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"age": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"salary": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	})
)
