package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtUserClaims struct {
	Name  string `json:"name"`
	jwt.StandardClaims
}

var userCount = 0
var JWT_SECRET = []byte("secret")


func registrateAnonym(c echo.Context) error {

	userCount++;
	// Set custom claims
	claims := &jwtUserClaims{
		"Jon Snow #" + fmt.Sprint(userCount),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
		"user": claims.Name,
	})

}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func userInfo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtUserClaims)
	name := claims.Name
	expired := claims.ExpiresAt

	return c.JSON(http.StatusOK, echo.Map{
		"user": name,
		"expiredAt": time.Unix(expired, 0),
	})
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// CORS userInfo
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))


	// Unauthenticated route
	e.GET("/", accessible)

	// Login route
	e.POST("/registrate_anonym", registrateAnonym)

	// Restricted group
	r := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &jwtUserClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user", userInfo)

	e.Logger.Fatal(e.Start(":3001"))
}