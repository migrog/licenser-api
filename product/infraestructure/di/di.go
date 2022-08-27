package di

import (
	"github.com/migrog/licenser-api/product/application/controller"
	"github.com/migrog/licenser-api/product/domain/interfaces"
	"github.com/migrog/licenser-api/product/domain/service"
	"github.com/migrog/licenser-api/product/infraestructure/database"
	"github.com/migrog/licenser-api/product/infraestructure/repository"
)

var (
	MongoConnection   database.IMongoConnection
	ProductRepository interfaces.IProductRepository
	ProductService    interfaces.IProductService
	ProductController *controller.ProductController
)

func SetupDependencyInjection(mongoConnection database.IMongoConnection) {
	MongoConnection = mongoConnection
	ProductRepository = repository.NewProductRepository(MongoConnection)
	ProductService = service.NewProductService(ProductRepository)
	ProductController = controller.NewProductController(ProductService)
}
