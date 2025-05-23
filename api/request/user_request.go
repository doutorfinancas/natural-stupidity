package request

// CreateUserRequest defines the payload to create a new user.
type CreateUserRequest struct {
	Name   string `json:"name" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	RoleID uint   `json:"role_id" binding:"required"`
	TeamID uint   `json:"team_id" binding:"required"`
}

// UpdateUserRequest defines the payload to update an existing user.
type UpdateUserRequest struct {
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty" binding:"omitempty,email"`
	RoleID *uint   `json:"role_id,omitempty"`
	TeamID *uint   `json:"team_id,omitempty"`
}
