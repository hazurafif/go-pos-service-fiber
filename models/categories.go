package models

type Categories struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(300)" json:"name"`
	CreatedAt string `gorm:"type:varchar(25)" json:"created_at"`
	UpdatedAt string `gorm:"type:varchar(25)" json:"updated_at"`
}
