package dto

type ProductResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewProductResponse(id string, name string) *ProductResponse {
	return &ProductResponse{
		Id:   id,
		Name: name,
	}
}
