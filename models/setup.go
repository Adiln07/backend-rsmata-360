package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("admin:S!MRSGos2@tcp(201.131.0.23:3306)/simata"))
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&Floor{})
	database.AutoMigrate(&Room{})

	DB = database

}