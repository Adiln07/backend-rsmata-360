package routers

import (
	mapController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func MapRoutes(router fiber.Router){
	router.Get("/floor-with-rooms", mapController.GetAllFloorsWithRooms)
	router.Get("/floor-with-rooms/:id", mapController.GetFLoorByIdWithRooms)
	router.Patch("/room-with-children/:id", mapController.GetRoomByIdWithChildren)
}