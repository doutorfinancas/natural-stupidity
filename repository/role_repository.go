package repository

import "github.com/doutorfinancas/natural-stupidity/repository/model"

// RoleRepository defines access to role records.
type RoleRepository interface {
	Create(role *model.Role) error
	GetByID(id uint) (*model.Role, error)
	List() ([]model.Role, error)
	Update(role *model.Role) error
	Delete(id uint) error
}
