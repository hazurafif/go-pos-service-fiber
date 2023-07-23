package models

import "time"

type Payment struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(300)" json:"name"`
	Type      string    `gorm:"type:varchar(300)" json:"type"`
	Logo      string    `gorm:"type:varchar(300)" json:"logo"`
	CreatedAt time.Time `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME" json:"updated_at"`
}
