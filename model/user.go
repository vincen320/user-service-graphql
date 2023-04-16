package model

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	cError "github.com/vincen320/user-service-graphql/helper/error"
	tokenModel "github.com/vincen320/user-service-graphql/model/token"
)

type (
	User struct {
		ID       int64   `json:"id"`
		Name     string  `json:"name"`
		Email    string  `json:"email"`
		Password string  `json:"-"`
		Age      int     `json:"age"`
		Address  string  `json:"address"`
		Salary   float64 `json:"salary"`
		Hobbies  []Hobby `json:"hobbies"`
	}

	UserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (u User) GenerateJWTToken() (tokenString string, err error) {
	var claims tokenModel.JWTClaims

	now := time.Now()
	claims.Id = fmt.Sprint(u.ID)
	claims.Issuer = "user-service-graphql"
	claims.IssuedAt = now.UTC().Unix()
	claims.ExpiresAt = now.Add(time.Hour * 1).UTC().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err = token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		err = cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		return
	}
	return
}
