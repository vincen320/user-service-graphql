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
	if user.Email = strings.TrimSpace(user.Email); user.Email == "" {
		return cError.New(http.StatusBadRequest, "email cannot be empty", "error email validation")
	}
	if user.Password = strings.TrimSpace(user.Password); user.Password == "" {
		return cError.New(http.StatusBadRequest, "password cannot be empty", "error password validation")
	}
	if len(user.Password) < 8 {
		return cError.New(http.StatusBadRequest, "password minumum has 8 character", "error password validation")
	}
	return nil
}

func ValidateUserLogin(user model.UserLogin) error {
	if user.Email = strings.TrimSpace(user.Email); user.Email == "" {
		return cError.New(http.StatusBadRequest, "email cannot be empty", "error email validation")
	}
	if user.Password = strings.TrimSpace(user.Password); user.Password == "" {
		return cError.New(http.StatusBadRequest, "password cannot be empty", "error password validation")
	}
	return nil
}
