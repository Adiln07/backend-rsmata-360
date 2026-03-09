package routers

import (
	hotspotInformationController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func HotspotInformationRoutes(router fiber.Router) {
	router.Get("/", hotspotInformationController.GetAllInformation)
	router.Get("/:id", hotspotInformationController.GetInformationById)
	router.Post("/", hotspotInformationController.CreateInformation)
	router.Patch("/:id", hotspotInformationController.UpdateInformation)
	router.Delete("/:id", hotspotInformationController.DeleteInformation)
}