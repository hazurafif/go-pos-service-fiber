package models

import "time"

type Product struct {
	Id          int64
	Sku         int64
	Name        string
	Stock       int64
	Price       int64
	Image       string
	Category_id int64
	Categories  Categories
	Created_at  time.Time
	Updated_at  time.Time
}
