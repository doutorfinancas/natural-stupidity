package service

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// VacationService encapsulates vacation-related business logic.
type VacationService interface {
	Create(vacation *model.Vacation) error
	GetByID(id uint) (*model.Vacation, error)
	List() ([]model.Vacation, error)
	ListByUser(userID uint) ([]model.Vacation, error)
	Update(vacation *model.Vacation) error
	Delete(id uint) error
}
