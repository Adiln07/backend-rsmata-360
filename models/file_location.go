package models

type FileLocation struct {
	Id   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255);unique" json:"name"`
	Path string `gorm:"type:varchar(255)" json:"path"`
}

func (FileLocation) TableName() string {
	return "file_location"
}
