package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
)

type hub struct {
	clients []*client
}

func newHub() *hub {
	return &hub{}
}

func (h *hub) run() {
	for {
		select {}
	}
}

// func (hub *WsHub) removeConn(clientId string) {
// 	hub.mu.Lock()
// 	defer hub.mu.Unlock()
// 	for i, conn := range hub.conns {
// 		if conn.clientId == clientId {
// 			conn.done()
// 			hub.conns = append(hub.conns[:i], hub.conns[i+1:]...)
// 			log.Infof("CLient %s disconnected", conn.clientId)
// 			go hub.logoutFn(clientId)
// 		}
// 	}
// }

func (h *hub) Broadcast(msg *serverMsg) {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to encode ServerMessage as JSON %s", err.Error())
	}
	preparedMsg, err := websocket.NewPreparedMessage(websocket.TextMessage, jsonBytes)
	if err != nil {
		log.Fatalf("Failed to create Prepared Websocket Message %s", err.Error())
	}
	for _, conn := range hub.conns {
		go func(c *wsConn, statuses ...model.ConnStatus) {
			if len(statuses) > 0 {
				if c.status.StatusNotIn(statuses) {
					return
				}
			}
			c.sendPrepared(preparedMsg)
		}(conn, statuses...)
	}
}
