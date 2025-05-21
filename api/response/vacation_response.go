package response

import "time"

// VacationResponse defines the JSON shape of a vacation.
type VacationResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
