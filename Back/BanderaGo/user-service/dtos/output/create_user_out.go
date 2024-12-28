package output

import "time"

type CreateUserOut struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}
