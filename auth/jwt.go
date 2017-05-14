package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

// jwtCustomClaims are custom claims extending default ones.
type JWTUserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var JWT_SECRET = []byte("secret")

var JWT_MIDDLEWARE_CONFIG = middleware.JWTConfig{
	Claims:      &JWTUserClaims{},
	SigningKey:  JWT_SECRET,
	TokenLookup: "query:token",
}
