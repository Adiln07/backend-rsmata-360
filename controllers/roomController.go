package controllers

import (
	"backend-rsmata-360/requests"
	"backend-rsmata-360/usecases"
	"backend-rsmata-360/validators"
	"fmt"
	"os"
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

func UpdateRoom2(c *fiber.Ctx) error{

	id := c.Query("id");

	if id == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"Query params 'id' required",
		})
	}

	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message": errConv.Error(),
		})
	}

	filesData, err := c.MultipartForm()

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":err.Error(),
		})
	}

	if filesData == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"No files uploaded",
		})
	}

	files, ok := filesData.File["image"]
	if !ok || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"failed",
			"message":"No file Uploaded",
		})
	}


	file := files[0]

	fileMeta := usecases.FileMeta{
		Filename: file.Filename,
		Size: file.Size,
		ContentType: file.Header.Get("Content-Type"),
	}

	result, err := usecases.UploadCase(fileMeta)

	if err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message": err.Error(),
			})
		}

		name := c.FormValue("name")
		pos_x_str := c.FormValue("pos_x")
		pos_y_str := c.FormValue("pos_y")
		statusStr := c.FormValue("status")

		pos_x ,err :=  strconv.ParseFloat(pos_x_str, 64)

		if err != nil{
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":"failed",
					"message":err.Error(),
				})
			}

			pos_y ,err :=  strconv.ParseFloat(pos_y_str, 64)

		if err != nil{
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":"failed",
					"message":err.Error(),
				})
			}

		status, err := strconv.Atoi(statusStr)

		if err != nil{
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"status":"failed",
					"message":err.Error(),
				})
			}

		updateRequest := requests.RoomUpdateRequest{
			Name: name,
			Image: result.Url,
			Pos_x: pos_x,
			Pos_y: pos_y,
			Status: status,
		}

		if err := validators.Validate.Struct(updateRequest); err != nil{
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

	if err := c.SaveFile(file, "." + result.Url); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"failed",
			"message":err.Error(),
		})
	}

	errr := usecases.UpdateRoom(updateRequest, convertInt)

	if errr != nil {
		if errr.Error() == "room not found"{
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":"failed",
				"message":"Room not found",
			})
		}
		if errr.Error() == "invalid room id"{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message":"Invalid room id",
			})
		}
		if errr.Error() == "no data to update"{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":"failed",
				"message":"No data to update",
			})
		}
		
			os.Remove("." + result.Url)
		
		
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": errr.Error(),
		})
		
		
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Room updated successfully",
	})

}

// update room with floor
func UpdateRoom(c *fiber.Ctx) error {

	id := c.Query("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Query params 'id' required",
		})
	}

	convertInt, errConv := strconv.Atoi(id)

	if errConv != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": errConv.Error(),
		})
	}

	filesData, err := c.MultipartForm()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	if filesData == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "No files uploaded",
		})
	}

	files, ok := filesData.File["image"]
	if !ok || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "No file Uploaded",
		})
	}

	name := c.FormValue("name")
	pos_x_str := c.FormValue("pos_x")
	pos_y_str := c.FormValue("pos_y")
	statusStr := c.FormValue("status")

	var pos_x float64
	var pos_y float64
	var status int

	if pos_x_str != "" {
    val, err := strconv.ParseFloat(pos_x_str, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": "pos_x must be a number",
        })
    }
    pos_x = val
}

if pos_y_str != "" {
    val, err := strconv.ParseFloat(pos_y_str, 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": "pos_y must be a number",
        })
    }
    pos_y = val
}

if statusStr != "" {
    val, err := strconv.Atoi(statusStr)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  "failed",
            "message": "status must be a number",
        })
    }
    status = val
}

	file := files[0]

	fileMeta := usecases.FileMeta{
		Filename:    file.Filename,
		Size:        file.Size,
		ContentType: file.Header.Get("Content-Type"),
	}

	result, err := usecases.UploadCase(fileMeta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	updateRequest := requests.RoomUpdateRequest{
		Name:   name,
		Image:  result.Url,
		Pos_x:  pos_x,
		Pos_y:  pos_y,
		Status: status,
	}

	if err := validators.Validate.Struct(updateRequest); err != nil {
		errors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0)

		for _, e := range errors {
			errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", e.Field(), e.Tag()))
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"errors": errorMessages,
		})
	}

	if err := c.SaveFile(file, "."+result.Url); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	errr := usecases.UpdateRoom(updateRequest, convertInt)

	if errr != nil {
		if errr.Error() == "room not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  "failed",
				"message": "Room not found",
			})
		}
		if errr.Error() == "invalid room id" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "failed",
				"message": "Invalid room id",
			})
		}
		if errr.Error() == "no data to update" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "failed",
				"message": "No data to update",
			})
		}

		os.Remove("." + result.Url)

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": errr.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Room updated successfully",
	})
}

func UpdateRoomEXP(c *fiber.Ctx) error{
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



