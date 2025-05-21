package response

// DirectionResponse defines the JSON shape of a direction.
type DirectionResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	DirectorID uint   `json:"director_id"`
}
