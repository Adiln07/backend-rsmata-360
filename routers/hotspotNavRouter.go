package routers

import (
	hotspotNavController "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func HotspotNavRoutes(router fiber.Router) {
	router.Get("/", hotspotNavController.GetAllNavigation)
	router.Get("/:id", hotspotNavController.GetNavigationById)
	router.Post("/", hotspotNavController.CreateNavigation)
	router.Patch("/:id", hotspotNavController.UpdateNavigation)
	router.Delete("/:id", hotspotNavController.DeleteNavigation)
}