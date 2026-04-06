package repositories

import (
	"backend-rsmata-360/models"
	"errors"

	"gorm.io/gorm"
)

func GetAllHotspotInformation() ([]models.HotspotInformation, error) {
	var hotspotInformations []models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information").Scan(&hotspotInformations).Error; err != nil {
		return []models.HotspotInformation{}, err
	}

	return hotspotInformations, nil
}

func GetHotspotInformationById(id int)(models.HotspotInformation, error){
	var hotspotInformation models.HotspotInformation

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, label, description, status FROM hotspot_information WHERE id = ?", id).Scan(&hotspotInformation).Error; err != nil{
		return models.HotspotInformation{}, err
	}

	if hotspotInformation.Id == 0{
		return models.HotspotInformation{},  errors.New("hotspot information not found")
	}

	return hotspotInformation, nil
}

func CreateInformation(hotspotInformation models.HotspotInformation) (models.HotspotInformation, error){

	
	if err := models.DB.Create(&hotspotInformation).Error; err != nil{
		return models.HotspotInformation{}, err
	}

	return hotspotInformation, nil	
}

func UpdateInformation( hotspotInformationUpdateRequest  map[string]interface{}, id int) (error){
	var hotspotInformation models.HotspotInformation

	if err := models.DB.First(&hotspotInformation, id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			 return errors.New("hotspot information not found")
		}
		return err
	}

	result := models.DB.Model(&hotspotInformation).Updates(hotspotInformationUpdateRequest)
	if result.Error != nil {
		 return result.Error
	}

	if result.RowsAffected == 0 {
		 return errors.New("cannot update the data")
	}

	return nil
}

func DeleteInformation(id int)(error){

	var hotspotInformation models.HotspotInformation

	result := models.DB.Delete(&hotspotInformation, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("hotspot information not found")
	}

	return nil
}
