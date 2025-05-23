package request

import "time"

// CreateVacationRequest defines the payload to create a new vacation.
type CreateVacationRequest struct {
	UserID    uint      `json:"user_id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
}

// UpdateVacationRequest defines the payload to update an existing vacation.
type UpdateVacationRequest struct {
	UserID    *uint      `json:"user_id,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}
