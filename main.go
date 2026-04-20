package main

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/routers"
	"backend-rsmata-360/validators"
	ws "backend-rsmata-360/websocket"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberWs "github.com/gofiber/websocket/v2"
)

func main() {
	validators.InitValidator()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods:"GET, POST, PUT, PATCH, DELETE",
	})) 

	models.ConnectDatabase()

	hub := ws.NewHub()
	ws.HubInstance = hub //Inject Global
	go hub.Run()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if fiberWs.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", ws.WebSocketHandler(hub))


	go func() {
		for {
			time.Sleep(5 * time.Second)
			hub.Broadcast <- []byte(`{"event":"test from server"}`)
		}
	}()

	api := app.Group("/api")

	app.Static("/uploads", "./uploads")

	api.Route("/floor", routers.FloorRoutes)
	api.Route("/upload", routers.UploadRoutes)
	api.Route("/room", routers.RoomRoutes)
	api.Route("/hotspot-information", routers.HotspotInformationRoutes)
	api.Route("/hotspot-navigation", routers.HotspotNavRoutes)
	api.Route("/map", routers.MapRoutes)  

	app.Listen(":8080")
}