package service

import (
	"github.com/migrog/licenser-api/license/domain/dto"
	"github.com/migrog/licenser-api/license/domain/entity"
	"github.com/migrog/licenser-api/license/domain/interfaces"
)

type LicenseService struct {
	LicenseRepository interfaces.ILicenseRepository
	UserFacade        interfaces.IUserFacade
	ProductFacade     interfaces.IProductFacade
}

func NewLicenseService(r interfaces.ILicenseRepository, fu interfaces.IUserFacade, fp interfaces.IProductFacade) interfaces.ILicenseService {
	return &LicenseService{LicenseRepository: r, UserFacade: fu, ProductFacade: fp}
}

func (s *LicenseService) CreateLicense(req *dto.CreateLicenseRequest) *dto.HttpResponse {
	u, err := s.UserFacade.UserByEmail(req.Email)
	if err != nil {
		return dto.BadRequest("License not registered")
	}

	p, err := s.ProductFacade.ProductById(req.ProductId)
	if err != nil {
		return dto.BadRequest("License not registered")
	}
	newLicense := entity.NewLicense()
	newLicense.NewLicenseKey()
	newLicense.AddUser(entity.NewUser(u.Email, u.Name))
	newLicense.AddProduct(entity.NewProduct(p.Id, p.Name))

	l, err := s.LicenseRepository.Create(newLicense)
	if err != nil {
		return dto.ServerError()
	}

	res := dto.NewLicenseResponse(l.Id.Hex(), l.Enabled, l.LicenseKey, l.Date)
	res.AddUser(l.User)
	res.AddProduct(l.Product)
	return dto.Created(res)
}
func (s *LicenseService) VerifyLicense(req *dto.VerifyLicenseRequest) *dto.HttpResponse {
	l, err := s.LicenseRepository.Verify(req)
	if err != nil {
		return dto.ServerError()
	}

	if l == nil {
		return dto.NotFound()
	}

	res := dto.NewLicenseResponse(l.Id.Hex(), l.Enabled, l.LicenseKey, l.Date)
	res.AddProduct(l.Product)
	res.AddUser(l.User)
	return dto.Ok(res)
}
func (s *LicenseService) EnableLicense(req *dto.EnableLicenseRequest) *dto.HttpResponse {
	l, err := s.LicenseRepository.Verify((*dto.VerifyLicenseRequest)(req))

	if err != nil {
		return dto.ServerError()
	}

	if l == nil {
		return dto.NotFound()
	}

	l.Enable()
	_, err = s.LicenseRepository.Update(l)
	if err != nil {
		return dto.ServerError()
	}

	res := dto.NewLicenseResponse(l.Id.Hex(), l.Enabled, l.LicenseKey, l.Date)
	res.AddProduct(l.Product)
	res.AddUser(l.User)

	return dto.Ok(res)
}

func (s *LicenseService) DisableLicense(req *dto.DisableLicenseRequest) *dto.HttpResponse {
	l, err := s.LicenseRepository.Verify((*dto.VerifyLicenseRequest)(req))

	if err != nil {
		return dto.ServerError()
	}

	if l == nil {
		return dto.NotFound()
	}

	l.Disable()
	_, err = s.LicenseRepository.Update(l)
	if err != nil {
		return dto.ServerError()
	}

	res := dto.NewLicenseResponse(l.Id.Hex(), l.Enabled, l.LicenseKey, l.Date)
	res.AddProduct(l.Product)
	res.AddUser(l.User)

	return dto.Ok(res)
}
