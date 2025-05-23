package request

// CreateRoleRequest defines the payload to create a new role.
type CreateRoleRequest struct {
	Name  string `json:"name" binding:"required"`
	Level int    `json:"level" binding:"required,min=1"`
}

// UpdateRoleRequest defines the payload to update an existing role.
type UpdateRoleRequest struct {
	Name  *string `json:"name,omitempty"`
	Level *int    `json:"level,omitempty,min=1"`
}
