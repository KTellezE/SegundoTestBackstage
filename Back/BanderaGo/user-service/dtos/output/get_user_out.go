package output

type GetUserOut struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}