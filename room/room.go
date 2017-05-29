package room

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/middleware"

	//"github.com/Vitaliy-Yarovuy/vw-server/auth"
	//"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/Vitaliy-Yarovuy/vw-server/auth"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	users = []string{"user1", "user 2"}
	
)

type game struct {
	author string
	time time.Time
}





func room(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"users": users,
		"games": []game {},
	})

}


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "pong !!34:2");
}

func wsUpgrade(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTUserClaims)
	name := claims.Name

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer func(){
		log.Printf("user %s - exit",name)
		ws.Close()
	} ()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Hello, %s !", name)))
		if err != nil {
			log.Print(err)
			return err
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Print(err)
			return err
		}
		fmt.Printf("MSG: %s\n", msg)
	}
}

func Linsten() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS userInfo
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Configure middleware with the custom claims type
	e.Use(middleware.JWTWithConfig(auth.JWT_MIDDLEWARE_CONFIG))

	e.GET("/", hello)
	e.GET("/ws", wsUpgrade)
	e.GET("/room", room)
	e.Logger.Fatal(e.Start(":3002"))
}
