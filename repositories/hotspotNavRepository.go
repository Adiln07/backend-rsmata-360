package repositories

import (
	"backend-rsmata-360/models"
	"errors"
)

func GetAllHotspotNav() ([]models.HotspotNav, error) {
	var hotspotNavigations []models.HotspotNav

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi").Scan(&hotspotNavigations).Error; err != nil{
		return []models.HotspotNav{}, err
	}
	return hotspotNavigations, nil
}

func GetHotspotNavById(id int)(models.HotspotNav, error){
	var hotspotNavigation models.HotspotNav

	if err := models.DB.Raw("SELECT id, room_id, yaw, pitch, description, target_room_label, target_room_id, status FROM hotspot_navigasi WHERE id = ?", id).Scan(&hotspotNavigation).Error; err != nil{
		return models.HotspotNav{}, err
	}

	if hotspotNavigation.Id == 0{
		return models.HotspotNav{},  errors.New("hotspot navigation not found")
	}

	return hotspotNavigation, nil
}

func CreateHotspotNav(hotspotNavigation models.HotspotNav) (models.HotspotNav, error){
	if err := models.DB.Create(&hotspotNavigation).Error; err != nil{
		return models.HotspotNav{}, err
	}
	return hotspotNavigation, nil
}

func UpdateHotspotNav( hotspotNavigationUpdateRequest  map[string]interface{}, id int) (error){

	result := models.DB.Model(&models.HotspotNav{}).Where("id = ?", id).Updates(hotspotNavigationUpdateRequest)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		 return errors.New("cannot update the data")
	}
	return nil
}

func DeleteHotspotNav(id int) (error){

	result := models.DB.Delete(&models.HotspotNav{}, id)

	if result.Error != nil{
		return result.Error
	}

	if result.RowsAffected == 0{
		return errors.New("hotspot navigation not found")
	}

	return nil
}

