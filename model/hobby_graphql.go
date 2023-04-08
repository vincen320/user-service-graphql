package model

import "github.com/graphql-go/graphql"

var (
	HobbyType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "hobby",
		Description: "object of hobby",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
)
