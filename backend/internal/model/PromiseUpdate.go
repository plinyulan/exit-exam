package model

import (
	"time"

	"gorm.io/gorm"
)

type PromiseUpdate struct {
	gorm.Model
	PromiseID uint      `gorm:"index;not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Note      string    `gorm:"not null"`

	Promise Promise
}

func (PromiseUpdate) TableName() string {
	return "promise_updates"
}
