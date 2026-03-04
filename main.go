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

	api := app.Group("/api");

	api.Route("/floor", routers.FloorRoutes)

	app.Listen(":8080")
}