package repository

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// TeamRepository defines access to team records.
type TeamRepository interface {
	Create(team *model.Team) error
	GetByID(id uint) (*model.Team, error)
	List() ([]model.Team, error)
	Update(team *model.Team) error
	Delete(id uint) error
}
