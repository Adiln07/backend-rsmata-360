package routers

import (
	uploadcontroller "backend-rsmata-360/controllers"

	"github.com/gofiber/fiber/v2"
)

func UploadRoutes(router fiber.Router) {
	router.Post("/", uploadcontroller.Upload)
}