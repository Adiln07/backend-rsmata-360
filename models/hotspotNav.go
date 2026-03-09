package models

type HotspotNav struct {
	Id                uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Room_Id           uint    `gorm:"column:room_id" json:"room_id"`
	Yaw               float64 `gorm:"column:yaw" json:"yaw"`
	Pitch             float64 `gorm:"column:pitch" json:"pitch"`
	Description       string  `gorm:"column:description" json:"description"`
	Target_Room_Label string  `gorm:"column:target_room_label" json:"target_room_label"`
	Target_Room_Id    uint    `gorm:"column:target_room_id" json:"target_room_id"`
	Status            int     `gorm:"column:status" json:"status"`
}

func (HotspotNav) TableName() string {
	return "hotspot_navigasi"
}