package usecase

import (
	"github.com/plinyulan/exit-exam/internal/model"
	"github.com/plinyulan/exit-exam/internal/services/repository"
)

type CampaignsUsecase interface {
	List() ([]model.Campaign, error)
}

type campUC struct {
	repo repository.CampaignsRepository
}

func NewCampaignsUsecase(r repository.CampaignsRepository) CampaignsUsecase {
	return &campUC{repo: r}
}

func (u *campUC) List() ([]model.Campaign, error) { return u.repo.List() }
