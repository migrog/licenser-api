package interfaces

import "github.com/migrog/licenser-api/license/domain/dto"

type IUserFacade interface {
	UserByEmail(email string) (user *dto.UserResponse, err error)
}
