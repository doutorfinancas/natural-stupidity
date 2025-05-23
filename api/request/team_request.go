package request

// CreateTeamRequest defines the payload to create a new team.
type CreateTeamRequest struct {
	Name        string `json:"name" binding:"required"`
	LeaderID    uint   `json:"leader_id" binding:"required"`
	DirectionID uint   `json:"direction_id" binding:"required"`
}

// UpdateTeamRequest defines the payload to update an existing team.
type UpdateTeamRequest struct {
	Name        *string `json:"name,omitempty"`
	LeaderID    *uint   `json:"leader_id,omitempty"`
	DirectionID *uint   `json:"direction_id,omitempty"`
}
