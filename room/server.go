package room

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/Vitaliy-Yarovuy/vw-server/auth"
	"time"
)


type game struct {
	author string
	time time.Time
}


func hello(c echo.Context) error {
	return c.String(http.StatusOK, "pong !!34:2");
}


// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.JWTUserClaims)
	name := claims.Name


	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	client := &Client{name: name, hub: hub, conn: conn, send: make(chan Command)}
	client.hub.register <- client
	go client.writePump()
	go func(){
		client.hub.broadcast <- enterRoomCommand(client.name)
	}()
	client.readPump()
	return nil
}


func Linsten() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	hub := newHub()
	go hub.run()


	// CORS userInfo
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Configure middleware with the custom claims type
	e.Use(middleware.JWTWithConfig(auth.JWT_MIDDLEWARE_CONFIG))

	e.GET("/", hello)
	e.GET("/ws",func(c echo.Context) error{
		return serveWs(hub, c)
	})
	e.GET("/room", func(c echo.Context) error{
		return c.JSON(http.StatusOK, echo.Map{
			"users": hub.getClients(),
			"games": []game {},
			"msgs": []string{},
		})
	})

	e.Logger.Fatal(e.Start(":3002"))
}
