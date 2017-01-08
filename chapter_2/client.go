package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザを表します
type client struct {
	socket   *websocket.Conn        // socketはこのクライアントのためのWebSocketです
	send     chan *message          // sendはメッセージが送られます
	room     *room                  // roomがこのクライアントが参加しているチャットルームです
	userData map[string]interface{} // userData はユーザに関する情報を保持します
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
