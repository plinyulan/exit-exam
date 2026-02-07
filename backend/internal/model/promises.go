package model

import (
	"time"

	"gorm.io/gorm"
)

type PromiseStatus string

const (
	PromiseNotStarted PromiseStatus = "not_started"
	PromiseInProgress PromiseStatus = "in_progress"
	PromiseFailed     PromiseStatus = "failed"
)

type Promise struct {
	gorm.Model
	PoliticianID uint          `gorm:"not null"`
	CampaignID   uint          `gorm:"not null"`
	Detail       string        `gorm:"not null"`
	AnnouncedAt  time.Time     `gorm:"not null"`
	Status       PromiseStatus `gorm:"not null"`

	Politician Politician
	Campaign   Campaign
	Updates    []PromiseUpdate
}

func (Promise) TableName() string {
	return "promises"
}
