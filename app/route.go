package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/vincen320/user-service-graphql/handler"
	"github.com/vincen320/user-service-graphql/repository"
	"github.com/vincen320/user-service-graphql/usecase"
	"golang.org/x/sync/errgroup"
)

func Run() {
	var (
		ctx, cancel    = context.WithCancel(context.Background())
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
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		cancel()
	}()
	e := echo.New()
	v1 := e.Group("v1")
	v1.POST("/graphql", graphqlHandler.GraphQL)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := e.Start(":3000")
		log.Println("successfully gracefully shutdown server")
		return err
	})
	eg.Go(func() error {
		<-egCtx.Done()
		err := db.Close()
		log.Println("DB Successfully closed")
		err = e.Shutdown(context.Background())
		log.Println("HTTP server finish graceful shutdown")
		return err
	})
	if err := eg.Wait(); err != nil {
		log.Println("Exit Reason", err)
	}
	fmt.Println("Clean up")
}
