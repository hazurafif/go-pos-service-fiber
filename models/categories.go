package models

import "time"

type Categories struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(300)" json:"name"`
	CreatedAt time.Time `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME" json:"updated_at"`
}
