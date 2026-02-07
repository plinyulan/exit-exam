package usecase

import (
	"errors"

	"time"

	"github.com/plinyulan/exit-exam/internal/model"
	"github.com/plinyulan/exit-exam/internal/services/repository"
)

type PromisesUsecase interface {
	ListAll() ([]model.Promise, error)
	GetDetail(id uint) (*model.Promise, error)
	ListByPolitician(politicianID uint) ([]model.Promise, error)
	AddUpdate(promiseID uint, updatedAt time.Time, note string) error
}

type promiseUC struct {
	repo repository.PromisesRepository
}

func NewPromisesUsecase(r repository.PromisesRepository) PromisesUsecase {
	return &promiseUC{repo: r}
}

func (u *promiseUC) ListAll() ([]model.Promise, error) {
	return u.repo.ListAllOrderedByAnnouncedAtDesc()
}

func (u *promiseUC) GetDetail(id uint) (*model.Promise, error) {
	return u.repo.GetDetail(id)
}

func (u *promiseUC) ListByPolitician(politicianID uint) ([]model.Promise, error) {
	return u.repo.ListByPolitician(politicianID)
}

// Business Rules:
// - Promise 1 ข้อ update ได้หลายครั้ง (OK)
// - ถ้าสถานะ failed ห้าม update เพิ่ม
func (u *promiseUC) AddUpdate(promiseID uint, updatedAt time.Time, note string) error {
	if note == "" {
		return errors.New("note is required")
	}
	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	p, err := u.repo.GetDetail(promiseID)
	if err != nil {
		return err
	}
	if p.Status == model.PromiseFailed {
		return errors.New("cannot update a failed promise")
	}

	up := &model.PromiseUpdate{
		PromiseID: promiseID,
		UpdatedAt: updatedAt,
		Note:      note,
	}
	return u.repo.CreateUpdate(up)
}
