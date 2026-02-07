package usecase

import (
	"github.com/plinyulan/exit-exam/internal/model"
	"github.com/plinyulan/exit-exam/internal/services/repository"
)

type PoliticiansUsecase interface {
	List() ([]model.Politician, error)
	Get(id uint) (*model.Politician, error)
}

type polUC struct {
	repo repository.PoliticiansRepository
}

func NewPoliticiansUsecase(r repository.PoliticiansRepository) PoliticiansUsecase {
	return &polUC{repo: r}
}

func (u *polUC) List() ([]model.Politician, error)      { return u.repo.List() }
func (u *polUC) Get(id uint) (*model.Politician, error) { return u.repo.GetByID(id) }
