package models

type Products struct {
	Id          int        `gorm:"primaryKey" json:"id"`
	Sku         int        `json:"sku"`
	Name        string     `json:"name"`
	Stock       int        `json:"stock"`
	Price       int        `json:"price"`
	Image       string     `json:"image"`
	Category_id int        `json:"category_id"`
	Categories  Categories `gorm:"foreignKey:Category_id" json:"categories"`
	CreatedAt   string     `gorm:"type:varchar(25)" json:"created_at"`
	UpdatedAt   string     `gorm:"type:varchar(25)" json:"updated_at"`
}
