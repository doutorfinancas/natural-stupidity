package repository

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// VacationRepository defines access to vacation records.
type VacationRepository interface {
	Create(vacation *model.Vacation) error
	GetByID(id uint) (*model.Vacation, error)
	List() ([]model.Vacation, error)
	ListByUser(userID uint) ([]model.Vacation, error)
	Update(vacation *model.Vacation) error
	Delete(id uint) error
}
