package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(localhost:3306)/golang"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Users{}, &Categories{}, &Order_product{}, &Orders{}, &Payment{}, &Products{})
	if err != nil {
		panic(err)
	}
	DB = db
}
