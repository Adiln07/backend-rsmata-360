package routers

import (
	roomController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoomRoutes(router fiber.Router) {
	router.Get("/", roomController.GetAllRoom)
	router.Get("/detail", roomController.GetRoomById)
	router.Post("/", roomController.CreateRoom)
	router.Patch("/", roomController.UpdateRoom)
	router.Delete("/", roomController.DeleteRoom)
} 