package model

import "gorm.io/gorm"

// User represents an application user.
// Implements core user profile data and relationships.
type User struct {
	gorm.Model
	Name         string `gorm:"size:100;not null" json:"name"`
	Email        string `gorm:"size:100;uniqueIndex;not null" json:"email"`
	RoleID       uint   `gorm:"not null" json:"role_id"`
	TeamID       uint   `gorm:"not null" json:"team_id"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
}
