package response

// RoleResponse defines the JSON shape of a role.
type RoleResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}
