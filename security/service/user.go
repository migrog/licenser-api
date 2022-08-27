package service

import "github.com/migrog/licenser-api/security/domain"

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(ur domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: ur,
	}
}

func (u *userService) GetById(id int64) (*domain.User, error) {
	return u.userRepo.GetById(id)

}

func (u *userService) GetByEmail(email string) (*domain.User, error) {
	return u.userRepo.GetByEmail(email)

}

func (u *userService) Create(m *domain.User) error {
	return u.userRepo.Create(m)
}
