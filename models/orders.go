package models

import "time"

type Orders struct {
	Id            int       `gorm:"primaryKey" json:"id"`
	UserID        int       `json:"user"`
	Users         Users     `gorm:"foreignKey:UserID" json:"users"`
	PaymentTypeID int       `json:"payment_type_id"`
	Payment       Payment   `gorm:"foreignKey:PaymentTypeID" json:"payment"`
	Name          string    `gorm:"type:varchar(300)" json:"name"`
	TotalPrice    int       `json:"total_price"`
	TotalPaid     int       `json:"total_paid"`
	ReceiptID     int       `json:"receipt_id"`
	CreatedAt     time.Time `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:DATETIME" json:"updated_at"`
}
