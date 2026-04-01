package controllers

import (
	"backend-rsmata-360/requests"
	"backend-rsmata-360/usecases"
	"backend-rsmata-360/validators"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// func Index(c *fiber.Ctx) error {
// 	var floors []models.Floor

// 	models.DB.Find(&floors)
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"floors": floors})
// }

// func Show(c *fiber.Ctx) error {
// 	var floor models.Floor
// 	id := c.Query("id")

// 	if id == ""{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Query parameter 'Id' is Required",
// 		})
// 	}

// 	convertInt, errConv := strconv.Atoi(id)

// 	if errConv != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": errConv.Error()})
// 	}

// 	if err := models.DB.First(&floor, convertInt).Error; err != nil {
// 		switch err {
// 		case gorm.ErrRecordNotFound:
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found "})

// 		default:
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
// 		}
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"floor": floor})
// }

func Index(c *fiber.Ctx) error{

	floors, err := usecases.GetAllFloors()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"data": floors,
		}) 
}

func Show(c *fiber.Ctx) error{
	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"Query Param 'id' is required",
		})
	}

	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message":errConv.Error(),
		})
	}

	floor, err := usecases.GetFloorById(convertInt)

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":"failed",
				"message": "Floor not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message":"Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"data": floor,
	})
}

// func Create(c *fiber.Ctx) error {
// 	var request requests.FloorCreateRequest
	
// 	if err := c.BodyParser(&request); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	if err := validators.Validate.Struct(request); err != nil{
// 		errors := err.(validator.ValidationErrors)
// 		errorMessages := make([]string, 0)

// 		for _, e := range errors{
// 			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
// 		}

// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"errors": errorMessages,
// 		})
// 	}

// 	floor := models.Floor{
// 		Name: request.Name,
// 		FloorPlan: request.FloorPlan,
// 		Status: request.Status,
// 	}

// 	result := models.DB.Create(&floor)

// 	if result.Error != nil{
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"floor": floor})
// }

func Create(c *fiber.Ctx) error{
	var floorRequest requests.FloorCreateRequest

	if err:= c.BodyParser(&floorRequest); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"success",
			"message":err.Error(),
		})
	}

	if err := validators.Validate.Struct(floorRequest); err != nil{
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)

		for _, e := range errors{
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"errors": errorMessages,
		})	
	}

	floor, err := usecases.CreateFloor(floorRequest)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": floor,
	})
}

// func Update(c *fiber.Ctx) error {
// 	id := c.Query("id") 

// 	if id == ""{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Query parameter 'Id' is Required",
// 		})
// 	}

// 	convertInt, errConv := strconv.Atoi(id)

// 	if errConv != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": errConv.Error()})
// 	}

// 	var request requests.FloorUpdateRequest
// 	if err := c.BodyParser(&request); err != nil{
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
// 	}

// 	updateMap :=map[string]interface{}{}

// 	if request.Name != nil{
// 		updateMap["name"] = *request.Name
// 	}

// 	if request.FloorPlan != nil {
// 		updateMap["floor_plan"] = *request.FloorPlan
// 	}
// 	if request.Status != nil {
// 		updateMap["status"] = *request.Status
// 	}

// 	if len(updateMap) == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No Data To Update"})
// 	}

// 	result :=models.DB.Model(&models.Floor{}).Where("id = ?", convertInt).Updates(updateMap)
	
// 	if result.Error != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
// 	}

// 	if result.RowsAffected == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Cannot Update the Data"})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Data Successfully Updated!!"})

// }

func Update(c *fiber.Ctx) error{
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "Query params 'id' required",
		})
	}

	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message": errConv.Error(),
		})
	}

	var updateRequest requests.FloorUpdateRequest

	if err := c.BodyParser(&updateRequest); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message": err.Error(),
		})
	}

	 err := usecases.UpdateFloor(updateRequest, convertInt)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"message":"data successfully updated",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Query("id") 

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message": "Query parameter 'Id' is Required",
		})
	}

	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": errConv.Error(),
		})
	}

	err := usecases.DeleteFloor(convertInt) 

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Data Completely Deleted!!",
	})
}

/*

! Kendala kode nya
* masih bingung cara untuk buat api yang bisa ada requirednya, jadi jika tidak user tidak isi salah satu data yang ada yang di isi tapi datanya required makan akan error 

*Pelajari semua kode pada bagian update nya karena sekarang kamu tidak mengerti 
*Pelajari semua kode pada bagian delete nanti karena sekarang kamu tidak mengerti

**/