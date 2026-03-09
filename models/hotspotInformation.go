package models

type HotspotInformation struct {
	Id          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Room_Id     uint    `gorm:"column:room_id" json:"room_id"`
	Yaw         float64 `gorm:"column:yaw" json:"yaw"`
	Pitch       float64 `gorm:"column:pitch" json:"pitch"`
	Label       string  `gorm:"column:label" json:"label"`
	Description string  `gorm:"column:description" json:"description"`
	Status      int     `gorm:"column:status" json:"status"`
}

func (HotspotInformation) TableName() string {
	return "hotspot_information"
}