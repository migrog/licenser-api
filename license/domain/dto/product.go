package dto

type ProductResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewProductResponse(id string, name string) *ProductResponse {
	return &ProductResponse{
		Id:   id,
		Name: name,
	}
}
