package service

import (
	"github.com/migrog/licenser-api/product/domain/dto"
	"github.com/migrog/licenser-api/product/domain/entity"
	"github.com/migrog/licenser-api/product/domain/interfaces"
)

type ProductService struct {
	ProductRepository interfaces.IProductRepository
}

func NewProductService(r interfaces.IProductRepository) interfaces.IProductService {
	return &ProductService{ProductRepository: r}
}

func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) *dto.HttpResponse {
	newProduct := entity.NewProduct(req.Name)
	p, err := s.ProductRepository.Create(newProduct)
	if err != nil {
		return dto.ServerError()
	}

	res := dto.NewProductResponse(p.Id.Hex(), p.Name)
	return dto.Created(res)
}
func (s *ProductService) ProductById(id string) *dto.HttpResponse {
	p, err := s.ProductRepository.GetById(id)
	if err != nil {
		return dto.ServerError()
	}

	if p == nil {
		return dto.NotFound()
	}

	res := dto.NewProductResponse(p.Id.Hex(), p.Name)
	return dto.Ok(res)

}
