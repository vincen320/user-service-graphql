package validator

import (
	"net/http"
	"strings"

	cError "github.com/vincen320/user-service-graphql/helper/error"
	"github.com/vincen320/user-service-graphql/model"
)

func ValidateCreateUser(user model.User) error {
	if user.Name = strings.TrimSpace(user.Name); user.Name == "" {
		return cError.New(http.StatusBadRequest, "name cannot be empty", "error name validation")
	}
	if user.Age < 0 {
		return cError.New(http.StatusBadRequest, "age cannot be below than zero", "error age validation")
	}
	if user.Address = strings.TrimSpace(user.Address); user.Address == "" {
		return cError.New(http.StatusBadRequest, "address cannot be empty", "error address validation")
	}
	if user.Salary < 0 {
		return cError.New(http.StatusBadRequest, "salary cannot be below than zero", "error salary validation")
	}
	return nil
}
