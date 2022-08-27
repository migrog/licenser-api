package di

import (
	"github.com/migrog/licenser-api/license/application/controller"
	"github.com/migrog/licenser-api/license/domain/interfaces"
	"github.com/migrog/licenser-api/license/domain/service"
	"github.com/migrog/licenser-api/license/infraestructure/database"
	"github.com/migrog/licenser-api/license/infraestructure/facade"
	"github.com/migrog/licenser-api/license/infraestructure/repository"
)

var (
	MongoConnection   database.IMongoConnection
	LicenseRepository interfaces.ILicenseRepository
	UserFacade        interfaces.IUserFacade
	ProductFacade     interfaces.IProductFacade
	LicenseService    interfaces.ILicenseService
	LicenseController *controller.LicenseController
)

func SetupDependencyInjection(mongoConnection database.IMongoConnection) {
	MongoConnection = mongoConnection
	LicenseRepository = repository.NewLicenseRepository(MongoConnection)
	UserFacade = facade.NewUserFacade()
	ProductFacade = facade.NewProductFacade()
	LicenseService = service.NewLicenseService(LicenseRepository, UserFacade, ProductFacade)
	LicenseController = controller.NewLicenseController(LicenseService)
}
