package dto

type UserResponse struct {
	Id    int64  `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func NewUserResponse(id int64, email, name string) *UserResponse {
	return &UserResponse{
		Id:    id,
		Email: email,
		Name:  name,
	}
}
