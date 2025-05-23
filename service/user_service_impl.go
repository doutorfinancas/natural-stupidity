package service

import (
	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
)

// userService is the implementation of UserService.
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(user *model.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) List() ([]model.User, error) {
	return s.repo.List()
}

func (s *userService) Update(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
