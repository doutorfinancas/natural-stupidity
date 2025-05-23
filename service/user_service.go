package service

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// UserService encapsulates user-related business logic.
type UserService interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	List() ([]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}
