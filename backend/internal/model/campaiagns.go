package model

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Year     int    `gorm:"not null"`
	District string `gorm:"not null"`
	Promises []Promise
}

func (Campaign) TableName() string {
	return "campaigns"
}
