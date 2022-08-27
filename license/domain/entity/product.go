package entity

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewProduct(id string, name string) *Product {
	return &Product{
		Id:   id,
		Name: name,
	}
}
