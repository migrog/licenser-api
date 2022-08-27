package entity

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUser(email, name string) *User {
	return &User{
		Email: email,
		Name:  name,
	}
}
