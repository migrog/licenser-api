package dto

import (
	"errors"
)

type DisableLicenseRequest struct {
	LicenseKey string `json:"license_key"`
	ProductId  string `json:"product_id"`
}

func (r *DisableLicenseRequest) IsValid() error {
	if r.LicenseKey == "" {
		return errors.New("license key is required")
	}

	if r.ProductId == "" {
		return errors.New("product id is required")
	}
	return nil
}
