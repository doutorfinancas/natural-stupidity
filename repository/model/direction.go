package model

import "gorm.io/gorm"

// Direction represents a collection of teams under a director.
type Direction struct {
	gorm.Model
	Name       string `gorm:"size:100;not null;unique" json:"name"`
	DirectorID uint   `gorm:"not null" json:"director_id"`
}
