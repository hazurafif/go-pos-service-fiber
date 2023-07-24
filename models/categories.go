package models

type Categories struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(300)" json:"name"`
	CreatedAt string `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt string `gorm:"type:DATETIME" json:"updated_at"`
}
