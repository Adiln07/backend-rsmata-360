package controllers

import (
	"backend-rsmata-360/requests"
	"backend-rsmata-360/usecases"
	"backend-rsmata-360/validators"
	"strconv"

	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllNavigation(c *fiber.Ctx) error {
	hotspotNavigations, err := usecases.GetAllHotspotNav()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": hotspotNavigations,
	})
}

func GetNavigationById(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "Query parameter 'Id' is Required",
		})
	}
	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errConv.Error()})
	}

	hotspotNavigation, err := usecases.GetHotspotNavById(convInt)

	if err != nil{
		if err.Error() == "hotspot navigation not found"{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Hotspot Navigation not found",
			})
		}
		if err.Error() == "invalid hotspot navigation id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Invalid hotspot navigation id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": hotspotNavigation,
	})
}


func CreateNavigation(c *fiber.Ctx) error {
	var createNavigationRequest requests.HotspotNavigationCreateRequest
	if err := c.BodyParser(&createNavigationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	if err := validators.Validate.Struct(createNavigationRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)

		for _, e := range errors {
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errorMessages,
		})
	}

	hotspotNavigation, err := usecases.CreateNavigation(createNavigationRequest)

	if err != nil {
		if err.Error() == "room not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Room not found",
			})
		}
		if err.Error() == "target room not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Target Room not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": hotspotNavigation,
	})
}

func UpdateNavigation(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "Query parameter 'Id' is Required",
		})
	}

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errConv.Error()})
	}

	var updateNavigationRequest requests.HotspotNavigationUpdateRequest

	if err := c.BodyParser(&updateNavigationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	err := usecases.UpdateNavigation(updateNavigationRequest, convInt)

	if err != nil {
		if err.Error() == "invalid hotspot navigation id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Invalid hotspot navigation id",
			})
		}
		if err.Error() == "cannot update the data" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Cannot update the data",
			})
		}
		if err.Error() == "room not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Room not found",
			})
		}
		if err.Error() == "target room not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Target Room not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Updated Navigation Successfully",
	})
}

func DeleteNavigation(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message": "Query parameter 'Id' is Required",
		})
	}

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errConv.Error(),
		})
	}

	err := usecases.DeleteNavigation(convInt)

	if err != nil{
		if err.Error() == "invalid hotspot navigation id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Invalid hotspot navigation id",
			})
		}
		if err.Error() == "hotspot navigation not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Hotspot Navigation not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Deleted Navigation Successfully",
	})
}