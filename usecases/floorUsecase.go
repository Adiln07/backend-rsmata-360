package usecases

import (
	"backend-rsmata-360/models"
	"backend-rsmata-360/repositories"
	"backend-rsmata-360/requests"
	"errors"
)

 
func GetAllFloors()([]models.Floor, error){
	floors , err := repositories.FindAllFloors()

	if err != nil {
		return []models.Floor{}, err
	}
	
	return floors, nil
}

func GetFloorById(id int) (models.Floor, error) {
	if id <= 0 {
		return models.Floor{}, errors.New("invalid floor id")
	}

	floor, err := repositories.FindFloorById(id)
	if err != nil {
		return models.Floor{}, err
	}

	return floor, nil
}

func CreateFloor(floorRequest requests.FloorCreateRequest)(models.Floor, error){

	floor := models.Floor{
	Name: floorRequest.Name,
	FloorPlan: floorRequest.FloorPlan,
	Status: floorRequest.Status,
}
	
	createdFloor, err := repositories.CreateFloor(floor)

	if err != nil{
		return models.Floor{}, err
	}

	return createdFloor, nil
}

func UpdateFloor(floorUpdateRequest requests.FloorUpdateRequest, id int)(error){

	if id <= 0 {
		return errors.New("invalid floor id")
	}

	updateMap := map[string]interface{}{}

	if floorUpdateRequest.Name != nil {
		updateMap["name"] = *floorUpdateRequest.Name
	}
	if floorUpdateRequest.FloorPlan != nil{
		updateMap["floor_plan"] = *floorUpdateRequest.FloorPlan
	}
	if floorUpdateRequest.Status != nil {
		updateMap["status"] = *floorUpdateRequest.Status
	}

	if len(updateMap) == 0{
		return errors.New("no data to update")
	}

	err := repositories.UpdateFloor(updateMap, id)

	return err
}

func DeleteFloor(id int)(error){
	
	if id <= 0 {
		return errors.New("invalid floor id")
	}

	err := repositories.DeleteFloor(id)

	return err
}