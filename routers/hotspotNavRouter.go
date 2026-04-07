package routers

import (
	hotspotNavController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func HotspotNavRoutes(router fiber.Router) {
	router.Get("/", hotspotNavController.GetAllNavigation)
	router.Get("/detail", hotspotNavController.GetNavigationById)
	router.Post("/", hotspotNavController.CreateNavigation)
	router.Patch("/", hotspotNavController.UpdateNavigation)
	router.Delete("/", hotspotNavController.DeleteNavigation)
}