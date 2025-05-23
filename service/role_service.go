package service

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// RoleService encapsulates role-related business logic.
type RoleService interface {
	Create(role *model.Role) error
	GetByID(id uint) (*model.Role, error)
	List() ([]model.Role, error)
	Update(role *model.Role) error
	Delete(id uint) error
}
