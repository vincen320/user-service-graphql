package app

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/vincen320/user-service-graphql/handler"
	"github.com/vincen320/user-service-graphql/repository"
	"github.com/vincen320/user-service-graphql/usecase"
)

func Run() {
	var (
		db             = NewDB()
		userRepository = repository.NewUserRepository(db)
		userUseCase    = usecase.NewUserUseCase(userRepository)
		userGQL        = handler.NewUserGQL(userUseCase)

		hobbyRepository = repository.NewHobbyRepository(db)
		hobbyUseCase    = usecase.NewHobbyUseCase(hobbyRepository)
		hobbyGQL        = handler.NewHobbyGQL(hobbyUseCase)

		queryType = graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user":  userGQL.GetUsers(),
				"hobby": hobbyGQL.GetHobbies(),
			},
		})

		mutationType = graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"user":  userGQL.CreateUser(),
				"login": userGQL.Login(),
			},
		})

		schema, _ = graphql.NewSchema(graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		})

		graphqlHandler = handler.NewGraphqlHandler(schema)
	)
	e := echo.New()
	v1 := e.Group("v1")
	v1.POST("/graphql", graphqlHandler.GraphQL)
	e.Start(":3000")
}
