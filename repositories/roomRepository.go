package repositories

import (
	"backend-rsmata-360/models"
	"errors"

	"gorm.io/gorm"
)

func GetAllRoom() ([]models.Room, error) {
	var rooms []models.Room

	err := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room").Scan(&rooms).Error

	if err != nil{
		return []models.Room{}, err
	}

	return rooms, nil
}

func GetRoomById(id int)(models.Room, error){
	var room models.Room

	err := models.DB.Raw("SELECT id, name, image, pos_x, pos_y, status FROM room WHERE id = ?", id).Scan(&room).Error;

	if err != nil{
		return models.Room{}, err
 	}

	if room.Id == 0{
		return models.Room{}, errors.New("room not found")
	}

	return room, nil
}

func CreateRoom(room models.Room, floorRoom models.FloorRoom)(models.Room, error){
	
	tx := models.DB.Begin()

	if tx.Error != nil {
		return models.Room{}, tx.Error
	}

	var floor models.Floor

	if err := tx.First(&floor, floorRoom.FloorId).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			tx.Rollback()
			return models.Room{}, errors.New("floor not found")
			}
			tx.Rollback()
			return models.Room{}, err
	}

	if err:= tx.Create(&room).Error; err != nil{
		tx.Rollback()
		return models.Room{}, err
	}

	floorRoom.RoomId = room.Id


	if err := tx.Create(&floorRoom).Error; err != nil{
		tx.Rollback()
		return models.Room{}, err
	}

	if err := tx.Commit().Error; err != nil{
		tx.Rollback()
		return models.Room{}, err
	}

	return room, nil
}

func UpdateRoom(roomUpdate models.Room, id int)(error){
	var room models.Room

	if err := models.DB.First(&room, id).Error ; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return errors.New("room not found")
		}
		return err
	}

	 result := models.DB.Model(&room).Updates(roomUpdate)

	 if result.Error != nil{
		return result.Error
	 }

	 if result.RowsAffected == 0 {
		return errors.New("cannot update room")
	 }

	return nil
}

func DeleteRoom(id int)(error){
	var room models.Room
	var floorRoom models.FloorRoom
	tx := models.DB.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.First(&room, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			tx.Rollback()
			return errors.New("room not found")
		}
		tx.Rollback()
		return err
	}

	if err := tx.Where("room_id = ?", id).Delete(&floorRoom).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&room).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil 
}