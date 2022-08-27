package interfaces

import "github.com/migrog/licenser-api/license/domain/dto"

type ILicenseService interface {
	CreateLicense(req *dto.CreateLicenseRequest) *dto.HttpResponse
	VerifyLicense(req *dto.VerifyLicenseRequest) *dto.HttpResponse
	EnableLicense(req *dto.EnableLicenseRequest) *dto.HttpResponse
	DisableLicense(req *dto.DisableLicenseRequest) *dto.HttpResponse
}
