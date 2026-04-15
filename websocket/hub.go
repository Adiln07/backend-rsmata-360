package websocket

import (
	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	clients 	map[*websocket.Conn]bool
	Broadcast 	chan[]byte
	register 	chan *websocket.Conn
	unregister 	chan *websocket.Conn
}

func NewHub() *Hub{
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte),
		register: make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run(){
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
		case conn := <- h.unregister:
			if _, ok := h.clients[conn]; ok{
				delete(h.clients, conn)
				conn.Close()
			}
		case message := <- h.Broadcast:
			for conn := range h.clients{
				err := conn.WriteMessage(websocket.TextMessage, message)

				if err != nil{
					conn.Close()
					delete(h.clients, conn)
				}
			}
		}
	}
}