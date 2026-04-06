package controllers

import (
	"backend-rsmata-360/requests"
	"backend-rsmata-360/usecases"
	"backend-rsmata-360/validators"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllInformation(c *fiber.Ctx) error{
	hotspotInformations, err := usecases.GetAllHotspotInformation()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": hotspotInformations,
	})
}

func GetInformationById(c *fiber.Ctx) error{
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "Query parameter 'Id' is Required",
		})
	}

	convInt, errorConv := strconv.Atoi(id)

	if errorConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errorConv.Error(),
		})
	}

	hotspotInformation, err := usecases.GetHotspotInformationById(convInt)

	if err != nil{
		if err.Error() == "hotspot information not found"{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Hotspot Information not found",
			})
		}

		if err.Error() == "invalid hotspot information id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Invalid hotspot information id",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": hotspotInformation,
	})
}


func CreateInformation(c *fiber.Ctx) error{

	var hotspotInformationCreateRequest requests.HotspotInformationCreateRequest

	if err := c.BodyParser(&hotspotInformationCreateRequest); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":err.Error(),
		})
	}

	if err := validators.Validate.Struct(hotspotInformationCreateRequest); err != nil {
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

	hotspotInformation, err := usecases.CreateInformation(hotspotInformationCreateRequest)

	if err != nil{
		if err.Error() == "room not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Room not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": hotspotInformation,
	})
}


func UpdateInformation(c *fiber.Ctx) error{
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message": "Query parameter 'Id' is Required",
		})
	}

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errConv.Error(),
		})
	}

	var hotspotInformationUpdateRequest requests.HotspotInformationUpdateRequest

	if err := c.BodyParser(&hotspotInformationUpdateRequest); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	err := usecases.UpdateInformation(hotspotInformationUpdateRequest, convInt)

	if err != nil{
		if err.Error() =="hotspot information not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Hotspot Information not found",
			})
		}
		if err.Error() == "invalid hotspot information id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Invalid hotspot information id",
			})
		}
		if err.Error() == "no data to update" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "No data to update",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Updated Successfully",
	})
}

func DeleteInformation(c *fiber.Ctx) error{
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
			"message": errConv.Error(),
		})
	}

	err := usecases.DeleteInfrmation(convInt)

	if err != nil{
		if err.Error() == "hotspot information not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "failed",
				"message": "Hotspot Information not found",
			})
		}
		if err.Error() == "invalid hotspot information id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "failed",
				"message": "Invalid hotspot information id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Deleted Successfully",
	})
}

// func DeleteInformationEXP(c *fiber.Ctx) error{
// 	var information models.HotspotInformation

// 	id := c.Query("id")

// 	if id == ""{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Query parameter 'Id' is Required",
// 		})
// 	}

// 	convInt, errConv := strconv.Atoi(id)

// 	if errConv != nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": errConv.Error()})
// 	}

// 	if err := models.DB.Delete(&information, convInt).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Deleted Successfully"})
// }
