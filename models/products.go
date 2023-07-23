package models

import "time"

type Products struct {
	Id          int        `gorm:"primaryKey" json:"id"`
	Sku         int        `json:"sku"`
	Name        string     `json:"name"`
	Stock       int        `json:"stock"`
	Price       int        `json:"price"`
	Image       string     `json:"image"`
	Category_id int        `json:"category_id"`
	Categories  Categories `gorm:"foreignKey:Category_id" json:"categories"`
	CreatedAt   time.Time  `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:DATETIME" json:"updated_at"`
}
