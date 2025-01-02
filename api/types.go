package api

type Profile struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Node struct {
	Name string `json:"name"`
}
