package models

type Floor struct {
	Id        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	FloorPlan string `gorm:"type:varchar(255)" json:"floor_plan"`
	Status    int    `gorm:"type:int" json:"status"`
}

func (Floor) TableName() string {
	return "floor"
}