package entities

import "time"

type Tipe struct {
	Id          uint
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}