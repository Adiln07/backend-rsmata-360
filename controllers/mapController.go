package controllers

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/responses"
	"backend-rsmata-360/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllFloorsWithRooms(c *fiber.Ctx) error {

	reponse, err := usecases.GetAllFloorsWithRooms()

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"failed",
			"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"data": reponse})
}

func GetFLoorByIdWithRooms(c *fiber.Ctx) error {

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

	response, err := usecases.GetFloorByIdWithRooms(convInt)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"data": response,
	}) 
}

func GetRoomByIdWithChildren(c *fiber.Ctx)error{
	id := c.Params("id")

	convInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	} 

	var room models.Room

	result := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room WHERE id = ?", convInt).Scan(&room)

	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": result.Error.Error()})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Room not Found"})
	}	

	var hotSpotInformation []models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information WHERE room_id = ?", convInt).Scan(&hotSpotInformation).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var hotspotNav []models.HotspotNav
	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi WHERE room_id = ?", convInt).Scan(&hotspotNav).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var infoReponses []responses.HotspotInformationResponse

	for _, info := range hotSpotInformation {
		infoReponses = append(infoReponses, responses.HotspotInformationResponse{
			Id :info.Id,
			Room_Id: info.Room_Id,
			Yaw: info.Yaw,
			Pitch: info.Pitch,
			Label: info.Label,
			Description: info.Description,
			Status: info.Status,
		})
	}

	var NavResponses []responses.HotspotNavigationResponse

	for _, nav := range hotspotNav{
		NavResponses = append(NavResponses, responses.HotspotNavigationResponse{
			Id:nav.Id,
			Room_Id: nav.Room_Id,
			Yaw: nav.Yaw,
			Pitch: nav.Pitch,
			Description: nav.Description,
			Target_Room_Label: nav.Target_Room_Label,
			Target_Room_Id: nav.Target_Room_Id,
			Status: nav.Status,
		})
	}

	responseRoom := responses.RoomWithChildrenResponse{
		Id: room.Id,
		Name: room.Name,
		Image: room.Image,
		Pos_x: room.Pos_x,
		Pos_y: room.Pos_y,
		Status: room.Status,
		HotspotInfo: infoReponses,
		HotspotNav: NavResponses,
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"room": responseRoom})
}
