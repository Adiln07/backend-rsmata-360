package main

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/routers"
	"backend-rsmata-360/validators"

	"github.com/gofiber/fiber/v2"
)

func main() {
	validators.InitValidator()

	app := fiber.New()

	models.ConnectDatabase()

	api := app.Group("/api")

	app.Static("/uploads", "./uploads")

	api.Route("/floor", routers.FloorRoutes)
	api.Route("/upload", routers.UploadRoutes)
	api.Route("/room", routers.RoomRoutes)
	api.Route("/hotspot-information", routers.HotspotInformationRoutes)
	api.Route("/hotspot-navigation", routers.HotspotNavRoutes)

	app.Listen(":8080")
}