package repositories

import (
	"backend-rsmata-360/models"
)

func MapGetAllFloors() ([]models.Floor, error) {
	var floors []models.Floor
	if err := models.DB.Find(&floors).Error; err != nil {
		return []models.Floor{}, err
	}

	return floors, nil
}

func MapGetFloorRoom()([]models.FloorRoom, error){
	var floorRooms []models.FloorRoom

	if err := models.DB.Find(&floorRooms).Error; err != nil{
		return []models.FloorRoom{}, err
	}
	return floorRooms, nil
}

func MapGetAllRooms()([]models.Room, error){
	var rooms []models.Room

	if err:= models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room").Scan(&rooms).Error; err != nil{
		return []models.Room{}, err
	}
	return rooms, nil
}

func MapGetAllHotspotInformation(roomIDs []uint) ([]models.HotspotInformation, error){
	var hotspotInformation []models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information WHERE room_id IN(?)", roomIDs).Scan(&hotspotInformation).Error; err != nil{
		return []models.HotspotInformation{}, err
	}
	return hotspotInformation, nil
}

func MapGetAllHotspotNavigation(roomIDs []uint) ([]models.HotspotNav, error){
	var hotspotNavigation []models.HotspotNav

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi WHERE room_id IN (?)", roomIDs).Scan(&hotspotNavigation).Error; err != nil{
		return []models.HotspotNav{}, err
	}
	return hotspotNavigation, nil
}

func MapGetFloorById(id int)(models.Floor, error){
	var floor models.Floor
	if err := models.DB.Where("id = ?", id).First(&floor).Error; err != nil{
		return models.Floor{}, err
	}
	return floor, nil
}

func MapGetFloorRoomByFloorId(floorId int)([]models.FloorRoom, error){
	var floorRooms []models.FloorRoom
	if err := models.DB.Where("floor_id = ?", floorId).Find(&floorRooms).Error; err != nil{
		return []models.FloorRoom{}, err
	}
	return floorRooms, nil
}

func MapGetAllRoomById(roomIDs []uint)([]models.Room, error){
	var rooms []models.Room

	if err := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room WHERE id IN(?)", roomIDs).Scan(&rooms).Error; err != nil{
		return []models.Room{}, err
	}
	return rooms, nil
}

