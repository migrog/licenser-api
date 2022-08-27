package facade

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/migrog/licenser-api/license/domain/dto"
	"github.com/migrog/licenser-api/license/domain/interfaces"
	"github.com/migrog/licenser-api/license/infraestructure/enviroment"
)

type UserFacade struct{}

func NewUserFacade() interfaces.IUserFacade {
	return &UserFacade{}
}

func (f *UserFacade) UserByEmail(email string) (user *dto.UserResponse, err error) {
	res, err := http.Get(fmt.Sprintf(enviroment.UserBaseURL+enviroment.UserRoute, email))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
