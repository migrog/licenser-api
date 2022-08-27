package domain

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

type UserService interface {
	GetByEmail(email string) (*User, error)
	GetById(id int64) (*User, error)
	Create(m *User) error
}

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	GetById(id int64) (*User, error)
	Create(m *User) error
}
