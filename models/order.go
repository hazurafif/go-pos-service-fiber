package models

import "time"

type Order struct {
	Id            int64     `gorm:"primaryKey" json:"id"`
	UserID        string    `json:"user_id"`
	User          Users     `gorm:"foreignKey:UserID" json:"user"`
	PaymentTypeID string    `json:"payment_type_id"`
	Payment       Payment   `gorm:"foreignKey:PaymentTypeID" json:"payment"`
	Name          string    `gorm:"type:varchar(300)" json:"name"`
	TotalPrice    int64     `json:"total_price"`
	TotalPaid     int64     `json:"total_paid"`
	ReceiptID     int64     `json:"receipt_id"`
	CreatedAt     time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp" json:"updated_at"`
}
