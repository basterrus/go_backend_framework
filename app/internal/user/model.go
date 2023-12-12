package user

type User struct {
	Name string `json:"name"`
}

type CreateUserDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateUserDTO struct {
}
