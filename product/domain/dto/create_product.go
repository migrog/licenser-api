package dto

import (
	"errors"
)

type CreateProductRequest struct {
	Name string `json:"name"`
}

func (r *CreateProductRequest) IsValid() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
