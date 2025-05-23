package service

import (
	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
)

// vacationService is the implementation of VacationService.
type vacationService struct {
	repo repository.VacationRepository
}

// NewVacationService creates a new VacationService.
func NewVacationService(repo repository.VacationRepository) VacationService {
	return &vacationService{repo: repo}
}

func (s *vacationService) Create(vacation *model.Vacation) error {
	return s.repo.Create(vacation)
}

func (s *vacationService) GetByID(id uint) (*model.Vacation, error) {
	return s.repo.GetByID(id)
}

func (s *vacationService) List() ([]model.Vacation, error) {
	return s.repo.List()
}

func (s *vacationService) ListByUser(userID uint) ([]model.Vacation, error) {
	return s.repo.ListByUser(userID)
}

func (s *vacationService) Update(vacation *model.Vacation) error {
	return s.repo.Update(vacation)
}

func (s *vacationService) Delete(id uint) error {
	return s.repo.Delete(id)
}
