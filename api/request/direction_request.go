package request

// CreateDirectionRequest defines the payload to create a new direction.
type CreateDirectionRequest struct {
	Name       string `json:"name" binding:"required"`
	DirectorID uint   `json:"director_id" binding:"required"`
}

// UpdateDirectionRequest defines the payload to update an existing direction.
type UpdateDirectionRequest struct {
	Name       *string `json:"name,omitempty"`
	DirectorID *uint   `json:"director_id,omitempty"`
}
