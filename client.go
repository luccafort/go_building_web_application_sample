package main

import (
	"github.com/gorilla/websocket"
)

type room struct {
	// forwardは他のクライアントに転送するためにメッセージを保持するチャンネル
	forward chan []byte
}

// clientはチャットを行っている1人のユーザを表します
type client struct {
	// socketはこのクライアントのためのWebSocketです
	socket *websocket.Conn
	// sendはメッセージが送られます
	send chan []byte
	// roomがこのクライアントが参加しているチャットルームです
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
