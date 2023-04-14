package token

import "github.com/dgrijalva/jwt-go"

type (
	JWTClaims struct {
		jwt.StandardClaims
	}
)
