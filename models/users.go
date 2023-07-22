package models

type Users struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(200)" json:"name"`
	Password  string `gorm:"-" json:"password"`
	CreatedAt string `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt string `gorm:"type:DATETIME" json:"updated_at"`
}
