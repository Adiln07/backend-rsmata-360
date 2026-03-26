package controllers

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/responses"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllFloorsWithRooms(c *fiber.Ctx) error {
	var floors []models.Floor
	if err := models.DB.Find(&floors).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if len(floors) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"floors": []responses.FloorRoomWithRoomsResponse{}})
	}

	var floorRooms []models.FloorRoom
	if err := models.DB.Find(&floorRooms).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	//* Map (floor_id -> room_id list)
	floorRoomMap := make(map[uint][]uint)
	for _, fr := range floorRooms{
		floorRoomMap[fr.FloorId] = append(floorRoomMap[fr.FloorId], fr.RoomId)
	}

 var rooms []models.Room
 if err:=  models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room").Scan(&rooms).Error; err != nil{
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
 }

//* Map cepat untuk mengambil room by id
roomMap := make(map[uint]models.Room)
roomIDs := []uint{}

for _, room := range rooms{
	roomMap[room.Id] = room
	roomIDs = append(roomIDs, room.Id)
}

// * Hotspot Information 
var hotSpotInformation []models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information WHERE room_id IN (?)", roomIDs).Scan(&hotSpotInformation).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()}) 
	}

	var hotspotNav []models.HotspotNav
	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi WHERE room_id IN (?)", roomIDs).Scan(&hotspotNav).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	mapInfo := make(map[uint][]models.HotspotInformation)
	mapNav := make(map[uint][]models.HotspotNav)

	for _, info := range hotSpotInformation {
		mapInfo[info.Room_Id] = append(mapInfo[info.Room_Id], info)
	}

	for _, nav := range hotspotNav {
		mapNav[nav.Room_Id] = append(mapNav[nav.Room_Id], nav)
	}

// *Response
var response []responses.FloorWithRoomsChildrenResponse

for _, floor := range floors{
	var roomList []responses.RoomWithChildrenResponse
	roomIDs := floorRoomMap[floor.Id]

	for _, rid := range roomIDs{
		if room, ok := roomMap[rid]; ok{

			// *convert hotspot info -> response type
			var infoResponse []responses.HotspotInformationResponse
			for _, info := range mapInfo[room.Id] {
				infoResponse =  append(infoResponse, responses.HotspotInformationResponse{
					Id: info.Id,
					Room_Id: info.Room_Id,
					Yaw: info.Yaw,
					Pitch: info.Pitch,
					Label: info.Label,
					Description: info.Description,
					Status: info.Status,
				})
			}

			// * convert hotspot Navigation -> response type
			var navResponse []responses.HotspotNavigationResponse
			for _, nav := range mapNav[room.Id]{
				navResponse = append(navResponse, responses.HotspotNavigationResponse{
					Id: nav.Id,
					Room_Id: nav.Room_Id,
					Yaw: nav.Yaw,
					Pitch: nav.Pitch,
					Description: nav.Description,
					Target_Room_Label: nav.Target_Room_Label,
					Target_Room_Id: nav.Target_Room_Id,
					Status: nav.Status,
				})
			}

			// * convert hotspot info & response tytpe
			roomList = append(roomList, responses.RoomWithChildrenResponse{
				Id:room.Id,
				Name: room.Name,
				Image: room.Image,
				Pos_x: room.Pos_x,
				Pos_y: room.Pos_y,
				Status: room.Status,
				HotspotInfo: infoResponse,
				HotspotNav: navResponse,
			})
		}
	}
	response = append(response, responses.FloorWithRoomsChildrenResponse{
		Id: floor.Id,
		Name: floor.Name,
		FloorPlan: floor.FloorPlan,
		Status: floor.Status,
		Rooms: roomList,
	})
}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"floors": response})
}

// func GetAllFloorsWithRooms(c *fiber.Ctx) error {
	
// 	return nil
// }

func GetFLoorByIdWithRooms(c *fiber.Ctx)error{

	// * Ambil ID floor dari parameter
	id := c.Params("id")
	convInt, errConv := strconv.Atoi(id)

	if errConv != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errConv.Error()})
	}

	// * Query apakah floornya ada by ID
	var floor models.Floor

	if err := models.DB.First(&floor, convInt).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var floorRoom []models.FloorRoom

	if err := models.DB.Where("floor_id = ?", convInt).Find(&floorRoom).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	roomIDs := []uint{}

	for _, fr := range floorRoom {
		roomIDs = append(roomIDs, fr.RoomId)
	}

	if len(roomIDs) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"floor" : responses.FloorWithRoomsChildrenResponse{}})
	}
	
	var rooms []models.Room
	if err := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room WHERE id IN(?)", roomIDs).Scan(&rooms).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

	// LANJUT CLUE 6 

	var hotspotInformation []models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information WHERE room_id IN(?)", roomIDs).Scan(&hotspotInformation).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	var hotspotNav []models.HotspotNav
	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi WHERE room_id IN(?)", roomIDs).Scan(&hotspotNav).Error; err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	mapInfo := make(map[uint][]models.HotspotInformation)
	mapNav := make(map[uint][]models.HotspotNav)

	for _, info := range hotspotInformation{
		mapInfo[info.Room_Id] = append(mapInfo[info.Room_Id], info)
	}

	for _, nav := range hotspotNav{
		mapNav[nav.Room_Id] = append(mapNav[nav.Room_Id], nav)
	}

	var roomResponses []responses.RoomWithChildrenResponse

	for _, room := range rooms{

		// * ambil children dari map
		infoList := mapInfo[room.Id]
		navList := mapNav[room.Id]

		// * konversi ke response struct
		var infoResponses []responses.HotspotInformationResponse
		for _, info := range infoList {
			infoResponses = append(infoResponses, responses.HotspotInformationResponse{
				Id: info.Id,
				Room_Id: info.Room_Id,
				Yaw : info.Yaw,
				Pitch: info.Pitch,
				Label : info.Label,
				Description: info.Description,
				Status: info.Status,
			})
		}

		var navResponses []responses.HotspotNavigationResponse
		for _, nav := range navList {
			navResponses = append(navResponses, responses.HotspotNavigationResponse{
				Id: nav.Id,
				Room_Id: nav.Room_Id,
				Yaw: nav.Yaw,
				Pitch:nav.Pitch,
				Description:nav.Description,
				Target_Room_Label: nav.Target_Room_Label,
				Target_Room_Id: nav.Target_Room_Id,
				Status: nav.Status,
			})
		}

		roomResponses = append(roomResponses, responses.RoomWithChildrenResponse{
			Id: room.Id,
			Name: room.Name,
			Image: room.Image,
			Pos_x: room.Pos_x,
			Pos_y: room.Pos_y,
			Status: room.Status,
			HotspotInfo: infoResponses,
			HotspotNav: navResponses,
		})

		
	}

	response := responses.FloorWithRoomsChildrenResponse{
		Id :floor.Id,
		Name: floor.Name,
		FloorPlan: floor.FloorPlan,
		Status: floor.Status,
		Rooms: roomResponses,
	}

	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"floor": response})
	
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
