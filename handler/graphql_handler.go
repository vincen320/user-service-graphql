package handler

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	cerror "github.com/vincen320/user-service-graphql/helper/error"
	"github.com/vincen320/user-service-graphql/model"
	responseModel "github.com/vincen320/user-service-graphql/model/response"
)

type graphQLHandler struct {
	graphqlSchema graphql.Schema
}

func NewGraphqlHandler(
	graphqlSchema graphql.Schema) *graphQLHandler {
	return &graphQLHandler{
		graphqlSchema: graphqlSchema,
	}
}

func (g *graphQLHandler) GraphQL(c echo.Context) (err error) {
	var request model.GraphQLRequest
	err = c.Bind(&request)
	if err != nil {
		log.Println(err)
		return responseModel.NewResponse(http.StatusBadRequest, "invalid payload").JSON(c.Response())
	}

	response := graphql.Do(graphql.Params{
		Context:        c.Request().Context(),
		Schema:         g.graphqlSchema,
		RequestString:  request.Query,
		OperationName:  request.OperationName,
		VariableValues: request.Variables,
	})

	if response.HasErrors() {
		if cerr, ok := cerror.ExtractCustomError(response.Errors[0].OriginalError()); ok {
			log.Println(cerr.GetActualError())
			return responseModel.NewResponse(cerr.GetCode(), cerr.GetErrorMessage()).JSON(c.Response())
		}
		log.Println(response.Errors)
		return responseModel.NewResponse(http.StatusBadRequest, response.Errors).JSON(c.Response())
	}
	return responseModel.NewResponse(http.StatusOK, response.Data).JSON(c.Response())
}
