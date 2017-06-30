package room

import (
	"github.com/satori/go.uuid"
)

const(
	ENTER_ROOM_ACTION string = "ENTER_ROOM_ACTION"
	LEAVE_ROOM_ACTION string = "LEAVE_ROOM_ACTION"
	USER_ENTERED_ROOM_ACTION string = "USER_ENTERED_ROOM_ACTION"
	USER_LEAVED_ROOM_ACTION string = "USER_LEAVED_ROOM_ACTION"
	SEND_MESSAGE_ACTION string = "SEND_MESSAGE_ACTION"
	RECEIVE_MESSAGE_ACTION string = "RECEIVE_MESSAGE_ACTION"
	CREATE_GAME_ACTION string = "CREATE_GAME_ACTION"
	JOIN_GAME_ACTION string = "JOIN_GAME_ACTION"
	lEAVE_GAME_ACTION string = "lEAVE_GAME_ACTION"
)

type Command struct {
	Id uuid.UUID `json:"id"`
	Type string `json:"type"`
	User string `json:"user"`
	Data string `json:"data"`
}



func enterRoomCommand(user string) Command{
	return Command{uuid.NewV1(),USER_ENTERED_ROOM_ACTION, user, ""}
}

func leaveRoomCommand(user string) Command{
	return Command{uuid.NewV1(), USER_LEAVED_ROOM_ACTION, user, ""}
}