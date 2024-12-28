package output

import "time"

type UpdateUserOut struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	UpdatedAt time.Time `json:"updated_at"`
}
