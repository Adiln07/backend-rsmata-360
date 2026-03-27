package controllers

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/requests"
	"backend-rsmata-360/validators"
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllRoom(c *fiber.Ctx) error {
	var rooms []models.Room

	
	if err := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room").Scan(&rooms).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"rooms": rooms})
}

func GetRoomById(c *fiber.Ctx) error{

	var room models.Room
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Query Params Id Is Required",
		})
	}

	convertInt , errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	if err := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room WHERE id = ?", convertInt).Scan(&room).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Data Not Found "})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"room": room})
}

func CreateRoom(c *fiber.Ctx) error{
	var request requests.RoomCreateRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validators.Validate.Struct(request); err != nil{
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)

		for _, e := range errors{
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Validation failed", "errors": errorMessages})
	}

	tx := models.DB.Begin()

	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": tx.Error.Error()})
	}

	var floor models.Floor

	if err := tx.First(&floor, request.FloorID).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			tx.Rollback()
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Floor not found"})
		}else{
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	room := models.Room{
		Name: request.Name,
		Image: request.Image,
		Pos_x: request.Pos_x,
		Pos_y: request.Pos_y,
		Status: request.Status,
	}

	if err := tx.Create(&room).Error; err != nil{
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	floorRoom := models.FloorRoom{
		FloorId: request.FloorID,
		RoomId: room.Id,
		Status: 1,
	}

	if err := tx.Create(&floorRoom).Error; err != nil{
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Room created successfully", "room": room})
}

func UpdateRoom(c *fiber.Ctx) error{

	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Query Params Id Is Required",
		})
	}

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": errConv.Error()})
	}

	var request requests.RoomUpdateRequest
	if err := c.BodyParser(&request); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	updateMap := make(map[string]interface{})

	if request.Name != nil {
		updateMap["name"] = *request.Name
	}
	if request.Image != nil {

		var oldRoom models.Room

		if err:= models.DB.First(&oldRoom, convInt).Error; err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		if oldRoom.Image != ""{
			oldPath := "." + oldRoom.Image
			_ = os.Remove(oldPath)
		}

		updateMap["image"] = *request.Image
	}
	if request.Pos_x != nil {
		updateMap["pos_x"] = *request.Pos_x
	}
	if request.Pos_y != nil {
		updateMap["pos_y"] = *request.Pos_y
	}
	if request.Status != nil {
		updateMap["status"] = *request.Status
	}

	if len(updateMap) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No data to update"})
	}

	var room models.Room

	if err := models.DB.First(&room, convInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Room not found"})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	result := models.DB.Model(&room).Updates(updateMap)
	
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": result.Error.Error()})
	}

	if result.RowsAffected == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Cannot Update the Data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Room updated successfully"})
}

func DeleteRoom(c *fiber.Ctx) error{
	var room models.Room

	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Query Params Id Is Required",
		})
	}
	convertInt, errConv := strconv.Atoi(id)
	
	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	tx := models.DB.Begin()

	defer tx.Rollback()

	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": tx.Error.Error()})
	}

	if err := tx.First(&room, convertInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Room not found"})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	}

	if err := tx.Where("room_id = ?", convertInt).Delete(&models.FloorRoom{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Delete(&room).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := tx.Commit().Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Room deleted successfully"})
}

