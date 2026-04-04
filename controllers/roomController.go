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

func GetAllRoom(c *fiber.Ctx) error{
	
	rooms, err := usecases.GetAllRoom()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"data": rooms, 
	})
}

func GetRoomById(c *fiber.Ctx) error{
	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"Query Params id is required",
		})
	}

	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":errConv.Error(),
		})
	}
	
	room, err := usecases.GetRoomById(convertInt)

	if err != nil {
		
		if err.Error() == "room not found"{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":"failed",
				"message":"Room not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message":"Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"data": room,
	})
}

func CreateRoom(c *fiber.Ctx) error {

	var roomCreateRequest requests.RoomCreateRequest
	if err := c.BodyParser(&roomCreateRequest); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":err.Error(),
		})
	}

	if err := validators.Validate.Struct(roomCreateRequest); err !=nil{
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)

		for _,e := range errors{
			errorMessages= append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"Validation Failed",
			"errors": errorMessages,
	})
	}

	room, err := usecases.CreateRoom(roomCreateRequest)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": room,
	})
}

func UpdateRoom(c *fiber.Ctx) error{
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

	var roomUpdateRequest requests.RoomUpdateRequest

	if err := c.BodyParser(&roomUpdateRequest); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	err := usecases.UpdateRoom(roomUpdateRequest, convertInt)

	if err != nil {
		if err.Error() == "room not found"{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":"failed",
				"message":"Room not found",
			})
		}
		if err.Error() == "invalid room id"{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message":"Invalid room id",
			})
		}
		if err.Error() == "no data to update"{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message":"No data to update",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Room updated successfully",
	})
}

func DeleteRoom(c *fiber.Ctx) error{

	id := c.Query("id")

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"Query Params Id Is Required",
		})
	}

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message":errConv.Error(),
		})
	}

	err := usecases.DeleteRoom(convInt)

	if err != nil{
		switch err.Error() {
		case "invalid room id":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message":"Invalid room id",
			})
		case "room not found":
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":"failed",
				"message":"Room not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":"failed",
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"message": "Room Deleted Successfully",
	})
}



