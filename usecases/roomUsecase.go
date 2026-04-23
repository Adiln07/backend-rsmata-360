package usecases

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/repositories"
	"backend-rsmata-360/requests"
	"errors"
	"os"
)

func GetAllRoom() ([]models.Room, error) {

	rooms, err := repositories.GetAllRoom()

	if err != nil{ 
		return []models.Room{}, err
	}
	
	return rooms, nil
}

func GetRoomById(id int)(models.Room, error){

	if id <= 0{
		return models.Room{},  errors.New("Invalid Room Id")
	}

	room, err := repositories.GetRoomById(id)

	if err != nil{
		return models.Room{}, err
	}

	return room, nil

}

func CreateRoom(roomCreateRequest requests.RoomCreateRequest)(models.Room, error){

	roomData := models.Room{
		Name: roomCreateRequest.Name,
		Image: roomCreateRequest.Image,
		Pos_x: roomCreateRequest.Pos_x,
		Pos_y: roomCreateRequest.Pos_y,
		Status: roomCreateRequest.Status,
	}

	floorRoom := models.FloorRoom{
		FloorId: roomCreateRequest.FloorID,
		Status: 1,
	}

	room, err := repositories.CreateRoom(roomData, floorRoom)

	if err != nil{
		return models.Room{}, err
	}

	return room, nil
}

func UpdateRoom(roomUpdateRequest requests.RoomUpdateRequest, id int)(error){
	if id <= 0{
		return errors.New("invalid room id")
	}

	oldRoom, errr := repositories.GetRoomById(id)

	if errr != nil{
		return errr
	}

	updatedata := models.Room{
		Name: roomUpdateRequest.Name,
		Image: roomUpdateRequest.Image,
		Pos_x: roomUpdateRequest.Pos_x,
		Pos_y: roomUpdateRequest.Pos_y,
		Status: roomUpdateRequest.Status,
	}

	err := repositories.UpdateRoom(updatedata, id)

	if err == nil && oldRoom.Image != "" && roomUpdateRequest.Image != oldRoom.Image{
			oldPath := "." + oldRoom.Image
			_ = os.Remove(oldPath)
		
	}

	return err
}

func DeleteRoom(id int)(error){

	if id <= 0{
		return errors.New("invalid room id")
	}
	err := repositories.DeleteRoom(id)
	return err
}