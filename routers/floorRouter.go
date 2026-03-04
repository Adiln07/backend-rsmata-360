package routers

import (
	floorcontroller "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func FloorRoutes(router fiber.Router) {
	router.Get("/", floorcontroller.Index)
	router.Get("/:id", floorcontroller.Show)
    router.Post("/", floorcontroller.Create)
    router.Patch("/:id", floorcontroller.Update)
    router.Delete("/:id", floorcontroller.Delete)
}