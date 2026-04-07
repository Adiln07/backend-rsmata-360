package usecases

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/repositories"
	"backend-rsmata-360/requests"
	"errors"
)

func GetAllHotspotInformation() ([]models.HotspotInformation, error) {
	hotspotInformations, err := repositories.GetAllHotspotInformation()

	if err != nil {
		return []models.HotspotInformation{}, err
	}
	return hotspotInformations, nil

}

func GetHotspotInformationById(id int) (models.HotspotInformation, error){

	if id <= 0{
		return models.HotspotInformation{}, errors.New("invalid hotspot information id")
	}

	hotspotInformation, err := repositories.GetHotspotInformationById(id)

	if err != nil{
		return models.HotspotInformation{}, err
	}

	return hotspotInformation, nil
}

func CreateInformation(hotspotInformationCreateRequest requests.HotspotInformationCreateRequest) (models.HotspotInformation, error){

	_, err := repositories.GetRoomById(int(hotspotInformationCreateRequest.Room_Id))

	if err != nil{
		return models.HotspotInformation{}, err
	}
	
	hotspotInformationData := models.HotspotInformation{
		Room_Id: hotspotInformationCreateRequest.Room_Id,
			Yaw: hotspotInformationCreateRequest.Yaw,
			Pitch: hotspotInformationCreateRequest.Pitch,
			Label: hotspotInformationCreateRequest.Label,
			Description: hotspotInformationCreateRequest.Description,
			Status: hotspotInformationCreateRequest.Status,
	}

	hotspotInformation, err := repositories.CreateInformation(hotspotInformationData)

	if err != nil{
		return models.HotspotInformation{}, err
	}
	return hotspotInformation, nil
}

func UpdateInformation(hotspotInformationUpdateRequest requests.HotspotInformationUpdateRequest, id int) (error){

	if id <= 0{
		return errors.New("invalid hotspot information id")
	}

	UpdateInformationRequest := make(map[string]interface{})

	if hotspotInformationUpdateRequest.Room_Id != nil {
		_, err := repositories.GetRoomById(int(*hotspotInformationUpdateRequest.Room_Id))
		if err != nil {
			return err
		}
		UpdateInformationRequest["room_id"] = *hotspotInformationUpdateRequest.Room_Id
	}
	if hotspotInformationUpdateRequest.Yaw != nil {
		UpdateInformationRequest["yaw"] = *hotspotInformationUpdateRequest.Yaw
	}
	if hotspotInformationUpdateRequest.Pitch != nil {
		UpdateInformationRequest["pitch"] = *hotspotInformationUpdateRequest.Pitch
	}
	if hotspotInformationUpdateRequest.Label != nil {
		UpdateInformationRequest["label"] = *hotspotInformationUpdateRequest.Label
	}
	if hotspotInformationUpdateRequest.Description != nil {
		UpdateInformationRequest["description"] = *hotspotInformationUpdateRequest.Description
	}
	if hotspotInformationUpdateRequest.Status != nil {
		UpdateInformationRequest["status"] = *hotspotInformationUpdateRequest.Status
	}

	if len(UpdateInformationRequest) == 0{
		return errors.New("no data to update")
	}

	err := repositories.UpdateInformation(UpdateInformationRequest, id)

	return err

}

func DeleteInformation(id int)(error){
	if id <= 0{
		return errors.New("invalid hotspot information id")
	}

	err := repositories.DeleteInformation(id)

	return err
}