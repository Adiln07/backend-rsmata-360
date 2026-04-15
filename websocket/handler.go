package websocket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func WebSocketHandler(hub *Hub) fiber.Handler{
	return websocket.New(func(c *websocket.Conn){

		hub.register <- c

		defer func() {
			hub.unregister <- c
		}()

		for {
			if _,_,err := c.ReadMessage(); err != nil{
				break 
			}
		}
	})
}