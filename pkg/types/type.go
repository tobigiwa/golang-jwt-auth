package types

import "github.com/golang-jwt/jwt/v4"

type UserCredentials struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ClientClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
