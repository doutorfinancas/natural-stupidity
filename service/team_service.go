package service

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// TeamService encapsulates team-related business logic.
type TeamService interface {
	Create(team *model.Team) error
	GetByID(id uint) (*model.Team, error)
	List() ([]model.Team, error)
	Update(team *model.Team) error
	Delete(id uint) error
}
