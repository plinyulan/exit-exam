package model

import "gorm.io/gorm"

type Politician struct {
	gorm.Model
	PoliticianCode string `gorm:"size:8;uniqueIndex;not null"` // 8 หลัก ตัวแรกไม่เป็น 0 จะตรวจใน usecase
	Name           string `gorm:"not null"`
	Party          string `gorm:"not null"`
	Promises       []Promise
}

func (Politician) TableName() string {
	return "politicians"
}
