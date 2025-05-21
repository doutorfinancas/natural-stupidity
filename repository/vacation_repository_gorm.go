package repository

import (
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"gorm.io/gorm"
)

// vacationRepository is a GORM-based implementation of VacationRepository.
type vacationRepository struct {
	db *gorm.DB
}

// NewVacationRepository creates a new VacationRepository using GORM.
func NewVacationRepository(db *gorm.DB) VacationRepository {
	return &vacationRepository{db: db}
}

func (r *vacationRepository) Create(vacation *model.Vacation) error {
	return r.db.Create(vacation).Error
}

func (r *vacationRepository) GetByID(id uint) (*model.Vacation, error) {
	var vac model.Vacation
	if err := r.db.First(&vac, id).Error; err != nil {
		return nil, err
	}
	return &vac, nil
}

func (r *vacationRepository) List() ([]model.Vacation, error) {
	var vacs []model.Vacation
	if err := r.db.Find(&vacs).Error; err != nil {
		return nil, err
	}
	return vacs, nil
}

func (r *vacationRepository) ListByUser(userID uint) ([]model.Vacation, error) {
	var vacs []model.Vacation
	if err := r.db.Where("user_id = ?", userID).Find(&vacs).Error; err != nil {
		return nil, err
	}
	return vacs, nil
}

func (r *vacationRepository) Update(vacation *model.Vacation) error {
	return r.db.Save(vacation).Error
}

func (r *vacationRepository) Delete(id uint) error {
	return r.db.Delete(&model.Vacation{}, id).Error
}
