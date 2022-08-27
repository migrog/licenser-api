package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func NewProduct(name string) *Product {
	return &Product{
		Id:   primitive.NewObjectID(),
		Name: name,
	}
}
