package entity

type User struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patrynomic string `json:"patrynomic"`
	Email      string `json:"email"`
	Phone      int    `json:"phone"`
	Password   string `json:"password"`
	Admin      bool   `json:"admin"`
}

func NewUser(name, surname, email, password string) *User {
	return &User{
		Name:     name,
		Surname:  surname,
		Password: password,
		Email:    email,
		Admin:    false,
	}
}
