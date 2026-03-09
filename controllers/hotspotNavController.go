package controllers

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/requests"
	"backend-rsmata-360/validators"
	"strconv"

	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllNavigation(c *fiber.Ctx) error {
	var hotspotNavigations []models.HotspotNav

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi").Scan(&hotspotNavigations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Navigations": hotspotNavigations})
}
func GetNavigationById(c *fiber.Ctx) error {
	var hotspotNavigation models.HotspotNav
	id := c.Params("id")
	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi WHERE id = ?", convInt).Scan(&hotspotNavigation).Error; err != nil{
		switch err {
			case gorm.ErrRecordNotFound:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found"})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": hotspotNavigation})
}
func CreateNavigation(c *fiber.Ctx) error {
	var request requests.HotspotNavigationCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validators.Validate.Struct(request); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)
		for _, e := range errors {
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}	
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": errorMessages})
 		}

		Navigation := models.HotspotNav{
			Room_Id: request.Room_Id,
			Yaw: request.Yaw,
			Pitch: request.Pitch,
			Description: request.Description,
			Target_Room_Label: request.Target_Room_Label,
			Target_Room_Id: request.Target_Room_Id,
			Status: request.Status,
		}

		if err := models.DB.Create(&Navigation).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Created Navigation Successfully"})
}
func UpdateNavigation(c *fiber.Ctx) error {
	id := c.Params("id")
	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	var request requests.HotspotNavigationUpdateRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updateMap := make(map[string]interface{})

	if request.Room_Id != nil {
		updateMap["room_id"] = *request.Room_Id
	}
	if request.Yaw != nil {
		updateMap["yaw"] = *request.Yaw
	}
	if request.Pitch != nil {
		updateMap["pitch"] = *request.Pitch
	}
	if request.Description != nil {
		updateMap["description"] = *request.Description
	}
	if request.Target_Room_Label != nil {
		updateMap["target_room_label"] = *request.Target_Room_Label
	}
	if request.Target_Room_Id != nil {
		updateMap["target_room_id"] = *request.Target_Room_Id
	}
	if request.Status != nil {
		updateMap["status"] = *request.Status
	}

	if len(updateMap) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No fields to update"})
	}

	var navigation models.HotspotNav

	result := models.DB.Model(&navigation).Where("id = ?", convInt).Updates(updateMap)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Updated Navigation Successfully"})
}
func DeleteNavigation(c *fiber.Ctx) error {
	var navigation models.HotspotNav
	id := c.Params("id")
	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})	
	}

	if err := models.DB.Delete(&navigation, convInt).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Deleted Navigation Successfully"})

}