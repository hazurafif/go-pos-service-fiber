package models

import "time"

type Payment struct {
	Id        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(300)" json:"name"`
	Type      string    `gorm:"type:varchar(300)" json:"type"`
	Logo      string    `gorm:"type:varchar(300)" json:"logo"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
