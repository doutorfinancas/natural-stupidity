package repository

import (
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"gorm.io/gorm"
)

// teamRepository is a GORM-based implementation of TeamRepository.
type teamRepository struct {
	db *gorm.DB
}

// NewTeamRepository creates a new TeamRepository using GORM.
func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) Create(team *model.Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepository) GetByID(id uint) (*model.Team, error) {
	var team model.Team
	if err := r.db.First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *teamRepository) List() ([]model.Team, error) {
	var teams []model.Team
	if err := r.db.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *teamRepository) Update(team *model.Team) error {
	return r.db.Save(team).Error
}

func (r *teamRepository) Delete(id uint) error {
	return r.db.Delete(&model.Team{}, id).Error
}
