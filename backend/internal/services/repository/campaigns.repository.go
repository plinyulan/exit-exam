package repository

import (
	"github.com/plinyulan/exit-exam/internal/model"
	"gorm.io/gorm"
)

type CampaignsRepository interface {
	List() ([]model.Campaign, error)
}

type campaignsRepo struct{ db *gorm.DB }

func NewCampaignsRepository(db *gorm.DB) CampaignsRepository {
	return &campaignsRepo{db: db}
}

func (r *campaignsRepo) List() ([]model.Campaign, error) {
	var items []model.Campaign
	return items, r.db.Order("year desc").Find(&items).Error
}
