package models

import "time"

type Users struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(200)" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(200)" json:"password"`
	CreatedAt time.Time `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:DATETIME" json:"updated_at"`
}
