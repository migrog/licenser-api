package dto

import (
	"errors"
)

type EnableLicenseRequest struct {
	LicenseKey string `json:"license_key"`
	ProductId  string `json:"product_id"`
}

func (r *EnableLicenseRequest) IsValid() error {
	if r.LicenseKey == "" {
		return errors.New("license key is required")
	}

	if r.ProductId == "" {
		return errors.New("product id is required")
	}
	return nil
}
