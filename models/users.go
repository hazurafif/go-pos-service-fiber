package models

type Users struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(200)" json:"name"`
	Password  string `gorm:"password()" json:"password"`
	CreatedAt string `gorm:"type:varchar(25)" json:"created_at"`
	UpdatedAt string `gorm:"type:varchar(25)" json:"updated_at"`
}
