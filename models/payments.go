package models

type Payment struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(300)" json:"name"`
	Type      string `gorm:"type:varchar(300)" json:"type"`
	Logo      string `gorm:"type:varchar(300)" json:"logo"`
	CreatedAt string `gorm:"type:varchar(25)" json:"created_at"`
	UpdatedAt string `gorm:"type:varchar(25)" json:"updated_at"`
}
