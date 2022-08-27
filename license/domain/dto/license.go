package dto

import (
	"time"

	"github.com/migrog/licenser-api/license/domain/entity"
)

type LicenseResponse struct {
	Id         string           `json:"id"`
	Enabled    bool             `json:"enabled"`
	LicenseKey string           `json:"license_key"`
	Product    *ProductResponse `json:"product,omitempty"`
	User       *UserResponse    `json:"user,omitempty"`
	Date       time.Time        `json:"date,omitempty"`
}

func NewLicenseResponse(id string, enabled bool, license_key string, date time.Time) *LicenseResponse {
	return &LicenseResponse{
		Id:         id,
		Enabled:    enabled,
		LicenseKey: license_key,
		Date:       date,
	}
}

func (r *LicenseResponse) AddProduct(m *entity.Product) {
	if m != nil {
		r.Product = &ProductResponse{
			Id:   m.Id,
			Name: m.Name,
		}
	}
}

func (r *LicenseResponse) AddUser(m *entity.User) {
	if m != nil {
		r.User = &UserResponse{
			Email: m.Email,
			Name:  m.Name,
		}
	}
}
