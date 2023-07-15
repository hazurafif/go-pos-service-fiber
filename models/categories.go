package models

type Categories struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(300)" json:"name"`
	Created_at string `gorm:"type:date" json:"created_at"`
	Updated_at string `gorm:"type:date" json:"updated_at"`
}
