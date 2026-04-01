package repositories

import (
	"backend-rsmata-360/models"
	"errors"
)

func FindAllFloors()([]models.Floor, error){
	var floors []models.Floor

	err := models.DB.Find(&floors).Error

	if err != nil{
		return []models.Floor{}, err
	}

	return floors, nil
}

func FindFloorById(id int)(models.Floor, error){
	var floor models.Floor

	err := models.DB.First(&floor, id).Error

	if err != nil{
		return models.Floor{}, err
	}
	return floor, nil
}

func CreateFloor(floor models.Floor)(models.Floor, error){
	
	err := models.DB.Create(&floor).Error

	if err != nil{
		return models.Floor{}, err
	}

	return floor, nil
}

func UpdateFloor(floorUpdate map[string]interface{}, id int)(error){

	var floor models.Floor

	result := models.DB.Model(&floor).Where("id= ?", id).Updates(floorUpdate)

	if result.Error != nil{
		return result.Error
	}

	if result.RowsAffected == 0{
		return errors.New("Floor not found")
	}

	return nil
}

func DeleteFloor(id int)(error){
	var floor models.Floor

	result := models.DB.Delete(&floor, id)

	if result.Error != nil{
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("Floor Not Found")
	}
	return nil
}