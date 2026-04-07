package usecases

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/repositories"
	"backend-rsmata-360/requests"
	"errors"
)

func GetAllHotspotNav() ([]models.HotspotNav, error) {
	hotspotNavigations, err := repositories.GetAllHotspotNav()

	if err != nil {
		return []models.HotspotNav{}, err
	}
	return hotspotNavigations, nil
}

func GetHotspotNavById(id int) (models.HotspotNav, error){
	if id <= 0{
		return models.HotspotNav{},  errors.New("invalid hotspot navigation id")
	}	

	hotspotNavigation, err := repositories.GetHotspotNavById(id)

	if err != nil{
		return models.HotspotNav{}, err
	}
	return hotspotNavigation, nil
}


func CreateNavigation(hotspotNavigation requests.HotspotNavigationCreateRequest) (models.HotspotNav, error){

	_, err := repositories.GetRoomById(int(hotspotNavigation.Room_Id))

	if err != nil{
		return models.HotspotNav{}, err
	}

	_, err = repositories.GetRoomById(int(hotspotNavigation.Target_Room_Id))

	if err != nil{
		return models.HotspotNav{}, errors.New("target room not found")
	}

	hotspotNavigationData := models.HotspotNav{
		Room_Id: hotspotNavigation.Room_Id,
			Yaw: hotspotNavigation.Yaw,
			Pitch: hotspotNavigation.Pitch,
			Description: hotspotNavigation.Description,
			Target_Room_Label: hotspotNavigation.Target_Room_Label,
			Target_Room_Id: hotspotNavigation.Target_Room_Id,
			Status: hotspotNavigation.Status,
	}

	hotspotNavigationResult, err := repositories.CreateHotspotNav(hotspotNavigationData)

	if err != nil{
		return models.HotspotNav{}, err
	}
	return hotspotNavigationResult, nil
}

func UpdateNavigation(hotspotNavigationUpdateRequest requests.HotspotNavigationUpdateRequest, id int) (error){
	if id <= 0{
		return errors.New("invalid hotspot navigation id")
	}

	UpdateNavigationRequest := make(map[string]interface{})

	if hotspotNavigationUpdateRequest.Room_Id != nil {
		_, err := repositories.GetRoomById(int(*hotspotNavigationUpdateRequest.Room_Id))
		if err != nil {
			return err
		}
		UpdateNavigationRequest["room_id"] = *hotspotNavigationUpdateRequest.Room_Id
	}
	if hotspotNavigationUpdateRequest.Yaw != nil {
		UpdateNavigationRequest["yaw"] = *hotspotNavigationUpdateRequest.Yaw
	}
	if hotspotNavigationUpdateRequest.Pitch != nil {
		UpdateNavigationRequest["pitch"] = *hotspotNavigationUpdateRequest.Pitch
	}
	if hotspotNavigationUpdateRequest.Description != nil {
		UpdateNavigationRequest["description"] = *hotspotNavigationUpdateRequest.Description
	}
	if hotspotNavigationUpdateRequest.Target_Room_Label != nil {
		UpdateNavigationRequest["target_room_label"] = *hotspotNavigationUpdateRequest.Target_Room_Label
	}
	if hotspotNavigationUpdateRequest.Target_Room_Id != nil {
		_, err := repositories.GetRoomById(int(*hotspotNavigationUpdateRequest.Target_Room_Id))
		if err != nil {
			return errors.New("target room not found")
		}
		UpdateNavigationRequest["target_room_id"] = *hotspotNavigationUpdateRequest.Target_Room_Id
	}
	if hotspotNavigationUpdateRequest.Status != nil {
		UpdateNavigationRequest["status"] = *hotspotNavigationUpdateRequest.Status
	}

	if len(UpdateNavigationRequest) == 0 {
		return errors.New("no data to update")
	}
	err := repositories.UpdateHotspotNav(UpdateNavigationRequest, id)

	return err

}

func DeleteNavigation(id int)(error){
	if id <= 0{
		return errors.New("invalid hotspot navigation id")
	}

	err := repositories.DeleteHotspotNav(id)

	return err
}