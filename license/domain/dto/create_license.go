package dto

import (
	"errors"
)

type CreateLicenseRequest struct {
	Email     string `json:"email"`
	ProductId string `json:"product_id"`
}

func (r *CreateLicenseRequest) IsValid() error {
	if r.Email == "" {
		return errors.New("email is required")
	}

	if r.ProductId == "" {
		return errors.New("product id is required")
	}
	return nil
}
