package models

type Order_product struct {
	Id          int      `gorm:"primaryKey" json:"id"`
	OrderId     int      `json:"order_id"`
	Orders      Orders   `gorm:"foreignKey:OrderId" json:"orders"`
	ProductId   int      `json:"product_id"`
	Products    Products `gorm:"foreignKey:ProductId" json:"products"`
	Qty         int      `json:"qty"`
	Total_price int      `json:"total_price"`
	CreatedAt   string   `gorm:"type:varchar(25)" json:"created_at"`
	UpdatedAt   string   `gorm:"type:varchar(25)" json:"updated_at"`
}
