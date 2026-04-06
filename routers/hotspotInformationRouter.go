package routers

import (
	hotspotInformationController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func HotspotInformationRoutes(router fiber.Router) {
	router.Get("/", hotspotInformationController.GetAllInformation)
	router.Get("/detail/", hotspotInformationController.GetInformationById)
	router.Post("/", hotspotInformationController.CreateInformation)
	router.Patch("/", hotspotInformationController.UpdateInformation)
	router.Delete("/", hotspotInformationController.DeleteInformation)
}