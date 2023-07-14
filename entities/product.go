package entities

import "time"

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Tipe    	Tipe
	Stock       int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}