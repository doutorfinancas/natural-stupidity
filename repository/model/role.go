package model

import "gorm.io/gorm"

// Role represents a user role with hierarchical level.
type Role struct {
	gorm.Model
	Name  string `gorm:"size:100;not null;unique" json:"name"`
	Level int    `gorm:"not null" json:"level"`
}
