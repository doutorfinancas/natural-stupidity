package service

import (
	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
)

// directionService is the implementation of DirectionService.
type directionService struct {
	repo repository.DirectionRepository
}

// NewDirectionService creates a new DirectionService.
func NewDirectionService(repo repository.DirectionRepository) DirectionService {
	return &directionService{repo: repo}
}

func (s *directionService) Create(direction *model.Direction) error {
	return s.repo.Create(direction)
}

func (s *directionService) GetByID(id uint) (*model.Direction, error) {
	return s.repo.GetByID(id)
}

func (s *directionService) List() ([]model.Direction, error) {
	return s.repo.List()
}

func (s *directionService) Update(direction *model.Direction) error {
	return s.repo.Update(direction)
}

func (s *directionService) Delete(id uint) error {
	return s.repo.Delete(id)
}
