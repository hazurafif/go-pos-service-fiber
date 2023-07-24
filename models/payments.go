package models

type Payment struct {
	Id        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(300)" json:"name"`
	Type      string `gorm:"type:varchar(300)" json:"type"`
	Logo      string `gorm:"type:varchar(300)" json:"logo"`
	CreatedAt string `gorm:"type:DATETIME" json:"created_at"`
	UpdatedAt string `gorm:"type:DATETIME" json:"updated_at"`
}
