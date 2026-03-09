package models

type FloorRoom struct {
	Id      uint `gorm:"primaryKey;autoIncrement" json:"id"`
	FloorId uint `gorm:"type:int" json:"floor_id"`
	RoomId  uint `gorm:"type:int" json:"room_id"`
	Status  int  `gorm:"type:int" json:"status"`
}

func (FloorRoom) TableName() string {
	return "floor_room"
}