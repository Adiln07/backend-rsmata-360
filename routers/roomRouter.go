package routers

import (
	roomController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoomRoutes(router fiber.Router) {
	router.Get("/", roomController.GetAllRoom)
	router.Get("/:id", roomController.GetRoomById)
	router.Post("/", roomController.CreateRoom)
	router.Patch("/:id", roomController.UpdateRoom)
	router.Delete("/:id", roomController.DeleteRoom)
}