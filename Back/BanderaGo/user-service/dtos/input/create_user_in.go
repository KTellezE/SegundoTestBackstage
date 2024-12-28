package input

type CreateUserIn struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
}
