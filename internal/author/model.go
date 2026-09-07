package author

import "time"

type Author struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Biography string    `json:"biography"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}