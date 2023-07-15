package models

import "time"

type Users struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(200)" json:"name"`
	Password   string    `gorm:"password()" json:"password"`
	Created_at time.Time `gorm:"type:timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp" json:"updated_at"`
}
