package repository

import (
	"gorm.io/gorm"

	"github.com/doutorfinancas/natural-stupidity/repository/model"
)

// roleRepository is a GORM-based implementation of RoleRepository.
type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a new RoleRepository using GORM.
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) List() ([]model.Role, error) {
	var roles []model.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepository) Update(role *model.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Role{}, id).Error
}
