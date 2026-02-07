package repository

import (
	"github.com/plinyulan/exit-exam/internal/model"
	"gorm.io/gorm"
)

type PromisesRepository interface {
	ListAllOrderedByAnnouncedAtDesc() ([]model.Promise, error)
	GetDetail(id uint) (*model.Promise, error)
	ListByPolitician(politicianID uint) ([]model.Promise, error)
	CreateUpdate(u *model.PromiseUpdate) error
}

type promisesRepo struct{ db *gorm.DB }

func NewPromisesRepository(db *gorm.DB) PromisesRepository {
	return &promisesRepo{db: db}
}

func (r *promisesRepo) ListAllOrderedByAnnouncedAtDesc() ([]model.Promise, error) {
	var items []model.Promise
	err := r.db.
		Preload("Politician").
		Preload("Campaign").
		Order("announced_at desc").
		Find(&items).Error
	return items, err
}

func (r *promisesRepo) GetDetail(id uint) (*model.Promise, error) {
	var p model.Promise
	err := r.db.
		Preload("Politician").
		Preload("Campaign").
		Preload("Updates", func(db *gorm.DB) *gorm.DB {
			return db.Order("updated_at desc")
		}).
		First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *promisesRepo) ListByPolitician(politicianID uint) ([]model.Promise, error) {
	var items []model.Promise
	err := r.db.
		Preload("Politician").
		Preload("Campaign").
		Where("politician_id = ?", politicianID).
		Order("announced_at desc").
		Find(&items).Error
	return items, err
}

func (r *promisesRepo) CreateUpdate(u *model.PromiseUpdate) error {
	return r.db.Create(u).Error
}
