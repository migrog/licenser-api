package database

import (
	"context"
	"log"

	"github.com/migrog/licenser-api/license/infraestructure/enviroment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IMongoConnection interface {
	Disconnect()
	GetCollection(dbName, collectionName string) (*mongo.Collection, error)
}

type MongoConnection struct {
	Client *mongo.Client
}

func NewConnection() IMongoConnection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(enviroment.MongoConnectionString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConnection{Client: client}
}

func (m *MongoConnection) GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	collection := m.Client.Database(dbName).Collection(collectionName)
	return collection, nil
}

func (m *MongoConnection) Disconnect() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}
