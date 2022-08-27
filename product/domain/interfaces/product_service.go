package interfaces

import "github.com/migrog/licenser-api/product/domain/dto"

type IProductService interface {
	CreateProduct(request *dto.CreateProductRequest) *dto.HttpResponse
	ProductById(id string) *dto.HttpResponse
}
