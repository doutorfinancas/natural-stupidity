package model

import "gorm.io/gorm"

// Team groups users under a leader and belongs to a direction.
type Team struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	LeaderID    uint   `gorm:"not null" json:"leader_id"`
	DirectionID uint   `gorm:"not null" json:"direction_id"`
}
