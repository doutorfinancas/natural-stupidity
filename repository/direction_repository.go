package repository

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// DirectionRepository defines access to direction records.
type DirectionRepository interface {
	Create(direction *model.Direction) error
	GetByID(id uint) (*model.Direction, error)
	List() ([]model.Direction, error)
	Update(direction *model.Direction) error
	Delete(id uint) error
}
