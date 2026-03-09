package controllers

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/requests"
	"backend-rsmata-360/validators"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllInformation(c *fiber.Ctx) error{
	var hotspotInformations []models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information").Scan(&hotspotInformations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"hotspot_informations": hotspotInformations})
}
func GetInformationById(c *fiber.Ctx) error{
	var hotspotInformation models.HotspotInformation
	id := c.Params("id")

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information WHERE id = ?", convInt).Scan(&hotspotInformation).Error; err != nil{
		switch err {
			case gorm.ErrRecordNotFound:
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found"})
			default:
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"hotspot_information": hotspotInformation})
}
func CreateInformation(c *fiber.Ctx) error{
	var request requests.HotspotInformationCreateRequest

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

		information := models.HotspotInformation{
			Room_Id: request.Room_Id,
			Yaw: request.Yaw,
			Pitch: request.Pitch,
			Label: request.Label,
			Description: request.Description,
			Status: request.Status,
			}

		if err := models.DB.Create(&information).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Created Create Successfully"})
}
func UpdateInformation(c *fiber.Ctx) error{

	id := c.Params("id")
	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}
	var request requests.HotspotInformationUpdateRequest

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
	if request.Label != nil {
		updateMap["label"] = *request.Label
	}
	if request.Description != nil {
		updateMap["description"] = *request.Description
	}
	if request.Status != nil {
		updateMap["status"] = *request.Status
	}

	if len(updateMap) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No fields to update"})
	}

	var information models.HotspotInformation

	if err := models.DB.First(&information, convInt).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found"})
		}else{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	result := models.DB.Model(&information).Updates(updateMap)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Cannot Update the Data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Updated Successfully"})
}
func DeleteInformation(c *fiber.Ctx) error{
	return nil
}
