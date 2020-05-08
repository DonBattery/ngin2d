package main

type messageType int

const (
	GameWorldDetails messageType = 1
	GameWorldUpdate  messageType = 2
)

type serverMsg struct {
	MsgType     messageType      `json:"msg_type"`
	GameWorld   *gameWorld       `json:"game_world"`
	GameObjects []gameObjectDump `json:"game_objects"`
}

func newServerMsg(msgType messageType, world *gameWorld, objects []gameObjectDump) *serverMsg {
	return &serverMsg{
		MsgType:     msgType,
		GameWorld:   world,
		GameObjects: objects,
	}
}

type gameWorld struct {
	BlockSize  int     `json:"block_size"`
	Background string  `json:"background"`
	Rows       [][]int `json:"rows"`
}

type gameObjectDump struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Color  string `json:"color"`
}
