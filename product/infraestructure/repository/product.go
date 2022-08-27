package repository

import (
	"context"
	"errors"
	"log"

	"github.com/migrog/licenser-api/product/domain/entity"
	"github.com/migrog/licenser-api/product/domain/interfaces"
	"github.com/migrog/licenser-api/product/infraestructure/database"
	"github.com/migrog/licenser-api/product/infraestructure/enviroment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository struct {
	Client database.IMongoConnection
}

func NewProductRepository(client database.IMongoConnection) interfaces.IProductRepository {
	return &ProductRepository{
		Client: client,
	}
}

func (r *ProductRepository) Create(p *entity.Product) (product *entity.Product, err error) {
	coll, err := r.Client.GetCollection(enviroment.ProductDatabase, enviroment.ProductCollection)
	if err != nil {
		log.Println(err)
	}
	_, err = coll.InsertOne(context.Background(), &p)
	if err != nil {
		log.Println(err)
		return nil, errors.New("An unexpected error occurred")
	}

	return p, nil
}

func (r *ProductRepository) GetById(id string) (product *entity.Product, err error) {
	coll, err := r.Client.GetCollection(enviroment.ProductDatabase, enviroment.ProductCollection)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	productId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("An unexpected error occurred")
	}

	if err = coll.FindOne(context.Background(), bson.M{"_id": productId}).Decode(&product); err != nil {
		log.Println(err)
		return nil, nil
	}
	return product, nil
}
