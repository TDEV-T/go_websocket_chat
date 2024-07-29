package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		//used to avoid concurrent writes on th websockte connection
		egress: make(chan []byte),
	}
}

func (c *Client) readMessage() {
	defer func() {
		//clearn connection

		c.manager.removeClient(c)
	}()
	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("err reading message : %v", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}
