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

type ProductFacade struct{}

func NewProductFacade() interfaces.IProductFacade {
	return &ProductFacade{}
}

func (f *ProductFacade) ProductById(id string) (product *dto.ProductResponse, err error) {
	res, err := http.Get(fmt.Sprintf(enviroment.ProductBaseURL+enviroment.ProductRoute, id))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&product); err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}
