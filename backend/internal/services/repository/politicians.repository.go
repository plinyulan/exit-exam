package repository

import (
	"github.com/plinyulan/exit-exam/internal/model"
	"gorm.io/gorm"
)

type PoliticiansRepository interface {
	List() ([]model.Politician, error)
	GetByID(id uint) (*model.Politician, error)
}

type politiciansRepo struct{ db *gorm.DB }

func NewPoliticiansRepository(db *gorm.DB) PoliticiansRepository {
	return &politiciansRepo{db: db}
}

func (r *politiciansRepo) List() ([]model.Politician, error) {
	var items []model.Politician
	return items, r.db.Order("created_at desc").Find(&items).Error
}

func (r *politiciansRepo) GetByID(id uint) (*model.Politician, error) {
	var p model.Politician
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
