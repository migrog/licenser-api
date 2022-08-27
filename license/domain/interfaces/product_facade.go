package interfaces

import "github.com/migrog/licenser-api/license/domain/dto"

type IProductFacade interface {
	ProductById(id string) (p *dto.ProductResponse, err error)
}
