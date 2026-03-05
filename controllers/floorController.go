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

func Index(c *fiber.Ctx) error {
	var floors []models.Floor

	models.DB.Find(&floors)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"floors": floors})
}

func Show(c *fiber.Ctx) error {
	var floor models.Floor
	id := c.Params("id")
	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	if err := models.DB.First(&floor, convertInt).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found "})

		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"floor": floor})
}


func Create(c *fiber.Ctx) error {
	var request requests.FloorCreateRequest
	
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validators.Validate.Struct(request); err != nil{
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)

		for _, e := range errors{
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errorMessages,
		})
	}

	floor := models.Floor{
		Name: request.Name,
		FloorPlan: request.FloorPlan,
		Status: request.Status,
	}

	result := models.DB.Create(&floor)

	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"floor": floor})
}

func Update(c *fiber.Ctx) error {
	id := c.Params("id") 
	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	var request requests.FloorUpdateRequest
	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updateMap :=map[string]interface{}{}

	if request.Name != nil{
		updateMap["name"] = *request.Name
	}

	if request.FloorPlan != nil {
		updateMap["floor_plan"] = *request.FloorPlan
	}
	if request.Status != nil {
		updateMap["status"] = *request.Status
	}

	if len(updateMap) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No Data To Update"})
	}

	result :=models.DB.Model(&models.Floor{}).Where("id = ?", convertInt).Updates(updateMap)
	
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot Update the Data"})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Data Successfully Updated!!"})

}

func Delete(c *fiber.Ctx) error {
	var floor models.Floor

	id := c.Params("id")
	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	if models.DB.Delete(&floor, convertInt).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot Delete the Data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Data Completely Deleted!!"})
}

/*

! Kendala kode nya
* masih bingung cara untuk buat api yang bisa ada requirednya, jadi jika tidak user tidak isi salah satu data yang ada yang di isi tapi datanya required makan akan error 

*Pelajari semua kode pada bagian update nya karena sekarang kamu tidak mengerti 
*Pelajari semua kode pada bagian delete nanti karena sekarang kamu tidak mengerti

**/