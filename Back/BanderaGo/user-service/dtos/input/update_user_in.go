package input

type UpdateUserIn struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}