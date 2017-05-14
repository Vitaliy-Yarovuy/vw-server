package main

import (
	"github.com/Vitaliy-Yarovuy/vw-server/auth"
	"github.com/Vitaliy-Yarovuy/vw-server/room"
)

func main() {

	go room.Linsten()
	go auth.Linsten()

	select {}

}
