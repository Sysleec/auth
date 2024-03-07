package model

import "github.com/dgrijalva/jwt-go"

const (
	ExamplePath = "/user/v1/create"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     int32  `json:"role"`
}

type LoginClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
