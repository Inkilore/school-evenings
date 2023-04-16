package gateway

type CreateUser struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Email    string `json:"email`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
