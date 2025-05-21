package response

// UserResponse defines the JSON shape of a user.
type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	TeamID uint   `json:"team_id"`
}
