package service

import (
	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
)

// roleService is the implementation of RoleService.
type roleService struct {
	repo repository.RoleRepository
}

// NewRoleService creates a new RoleService.
func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) Create(role *model.Role) error {
	return s.repo.Create(role)
}

func (s *roleService) GetByID(id uint) (*model.Role, error) {
	return s.repo.GetByID(id)
}

func (s *roleService) List() ([]model.Role, error) {
	return s.repo.List()
}

func (s *roleService) Update(role *model.Role) error {
	return s.repo.Update(role)
}

func (s *roleService) Delete(id uint) error {
	return s.repo.Delete(id)
}
