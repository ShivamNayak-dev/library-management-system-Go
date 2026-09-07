package author

type CreateAuthorRequest struct {
	Name      string `json:"name"`
	Biography string `json:"biography"`
}