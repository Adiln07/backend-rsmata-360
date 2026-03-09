package models

type Room struct {
	Id     uint    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name   string  `gorm:"column:name" json:"name"`
	Image  string  `gorm:"column:image" json:"image"`
	Pos_x  float64 `gorm:"column:pos_x" json:"pos_x"`
	Pos_y  float64 `gorm:"column:pos_y" json:"pos_y"`
	Status int     `gorm:"column:status" json:"status"`
}

func (Room) TableName() string {
	return "room"
}