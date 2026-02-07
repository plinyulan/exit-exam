package usecase

import "github.com/plinyulan/exit-exam/internal/services/repository"

type CatUsecase interface {
	GetCatsUsecase() []string
}

type catUsecase struct {
	repo repository.CatRepository
}

func NewCatUsecase(repo repository.CatRepository) CatUsecase {
	return &catUsecase{repo: repo}
}

func (u *catUsecase) GetCatsUsecase() []string {
	response := u.repo.GetCats()
	return response
}
