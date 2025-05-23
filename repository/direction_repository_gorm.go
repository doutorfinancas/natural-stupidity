package repository

import (
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"gorm.io/gorm"
)

// directionRepository is a GORM-based implementation of DirectionRepository.
type directionRepository struct {
	db *gorm.DB
}

// NewDirectionRepository creates a new DirectionRepository using GORM.
func NewDirectionRepository(db *gorm.DB) DirectionRepository {
	return &directionRepository{db: db}
}

func (r *directionRepository) Create(direction *model.Direction) error {
	return r.db.Create(direction).Error
}

func (r *directionRepository) GetByID(id uint) (*model.Direction, error) {
	var dir model.Direction
	if err := r.db.First(&dir, id).Error; err != nil {
		return nil, err
	}
	return &dir, nil
}

func (r *directionRepository) List() ([]model.Direction, error) {
	var dirs []model.Direction
	if err := r.db.Find(&dirs).Error; err != nil {
		return nil, err
	}
	return dirs, nil
}

func (r *directionRepository) Update(direction *model.Direction) error {
	return r.db.Save(direction).Error
}

func (r *directionRepository) Delete(id uint) error {
	return r.db.Delete(&model.Direction{}, id).Error
}
