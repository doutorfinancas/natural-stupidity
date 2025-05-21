package service

import (
	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
)

// teamService is the implementation of TeamService.
type teamService struct {
	repo repository.TeamRepository
}

// NewTeamService creates a new TeamService.
func NewTeamService(repo repository.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) Create(team *model.Team) error {
	return s.repo.Create(team)
}

func (s *teamService) GetByID(id uint) (*model.Team, error) {
	return s.repo.GetByID(id)
}

func (s *teamService) List() ([]model.Team, error) {
	return s.repo.List()
}

func (s *teamService) Update(team *model.Team) error {
	return s.repo.Update(team)
}

func (s *teamService) Delete(id uint) error {
	return s.repo.Delete(id)
}
