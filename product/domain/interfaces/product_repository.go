package interfaces

import "github.com/migrog/licenser-api/product/domain/entity"

type IProductRepository interface {
	Create(p *entity.Product) (product *entity.Product, err error)
	GetById(id string) (product *entity.Product, err error)
}
