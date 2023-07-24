package entities

import "time"

type Status string

const (
	Ready Status = "ready"
	TidakReady Status = "tidakready"
)

type Product struct {
	Id          uint
	Name        string
	Category    Category
	Tipe    	Tipe
	Brand    	Brand
	Status		Status `gorm:"type:enum('ready', 'tidakready')"`
	Stock       int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	No 			int
}