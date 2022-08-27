package repository

import (
	"context"
	"errors"
	"log"

	"github.com/migrog/licenser-api/license/domain/dto"
	"github.com/migrog/licenser-api/license/domain/entity"
	"github.com/migrog/licenser-api/license/domain/interfaces"
	"github.com/migrog/licenser-api/license/infraestructure/database"
	"github.com/migrog/licenser-api/license/infraestructure/enviroment"
	"go.mongodb.org/mongo-driver/bson"
)

type LicenseRepository struct {
	Client database.IMongoConnection
}

func NewLicenseRepository(client database.IMongoConnection) interfaces.ILicenseRepository {
	return &LicenseRepository{
		Client: client,
	}
}

func (r *LicenseRepository) Create(l *entity.License) (license *entity.License, err error) {
	coll, err := r.Client.GetCollection(enviroment.LicenseDatabase, enviroment.LicenseCollection)
	if err != nil {
		log.Println(err)
	}
	_, err = coll.InsertOne(context.Background(), &l)
	if err != nil {
		log.Println(err)
		return nil, errors.New("an unexpected error occurred")
	}

	return l, nil
}

func (r *LicenseRepository) Verify(req *dto.VerifyLicenseRequest) (license *entity.License, err error) {
	coll, err := r.Client.GetCollection(enviroment.LicenseDatabase, enviroment.LicenseCollection)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err != nil {
		log.Println(err)
		return nil, errors.New("an unexpected error occurred")
	}

	if err = coll.FindOne(context.Background(), bson.M{"license_key": req.LicenseKey, "product.id": req.ProductId}).Decode(&license); err != nil {
		log.Println(err)
		return nil, nil
	}
	return license, nil
}

func (r *LicenseRepository) Update(l *entity.License) (license *entity.License, err error) {
	coll, err := r.Client.GetCollection(enviroment.LicenseDatabase, enviroment.LicenseCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("an unexpected error occurred")
	}
	res, err := coll.UpdateOne(context.Background(), bson.M{"_id": l.Id}, bson.M{"$set": l})
	if err != nil {
		log.Println(err)
		return nil, errors.New("an unexpected error occurred")
	}

	if res.ModifiedCount == 0 {
		log.Println(err)
		return nil, errors.New("an unexpected error occurred")
	}

	return l, nil
}
