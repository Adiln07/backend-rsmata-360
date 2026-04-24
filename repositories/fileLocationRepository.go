package repositories

import (
	"backend-rsmata-360/models"
	"errors"

	"gorm.io/gorm"
)

func GetFileLocationByName(name string) (models.FileLocation, error){

	var fileLocation models.FileLocation

	err := models.DB.Where("name = ?", name).First(&fileLocation).Error

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return fileLocation, errors.New("file location not found")
		}
		return fileLocation, err
	}
	return fileLocation, nil
}