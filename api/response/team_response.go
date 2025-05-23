package response

// TeamResponse defines the JSON shape of a team.
type TeamResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	LeaderID    uint   `json:"leader_id"`
	DirectionID uint   `json:"direction_id"`
}
