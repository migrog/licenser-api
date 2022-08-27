package mysql

import (
	"github.com/migrog/licenser-api/security/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() domain.UserRepository {
	return &UserRepository{Conn}
}

func (m *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := m.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (m *UserRepository) GetById(id int64) (*domain.User, error) {
	var user domain.User

	err := m.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserRepository) Create(user *domain.User) error {
	return m.DB.Create(&user).Error

}
