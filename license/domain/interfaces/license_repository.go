package interfaces

import (
	"github.com/migrog/licenser-api/license/domain/dto"
	"github.com/migrog/licenser-api/license/domain/entity"
)

type ILicenseRepository interface {
	Create(l *entity.License) (license *entity.License, err error)
	Verify(req *dto.VerifyLicenseRequest) (license *entity.License, err error)
	Update(l *entity.License) (license *entity.License, err error)
}
