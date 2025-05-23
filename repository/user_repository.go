package repository

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// UserRepository defines access to user records.
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	List() ([]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
	FindByEmail(email string) (*model.User, error)
}
