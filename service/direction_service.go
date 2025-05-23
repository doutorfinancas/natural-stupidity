package service

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// DirectionService encapsulates direction-related business logic.
type DirectionService interface {
	Create(direction *model.Direction) error
	GetByID(id uint) (*model.Direction, error)
	List() ([]model.Direction, error)
	Update(direction *model.Direction) error
	Delete(id uint) error
}
