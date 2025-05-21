package model

import (
	"time"

	"gorm.io/gorm"
)

// Vacation represents a user's time off.
type Vacation struct {
	gorm.Model
	UserID    uint      `gorm:"not null" json:"user_id"`
	StartDate time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time `gorm:"type:date;not null" json:"end_date"`
}
