package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var once sync.Once
var Client *mongo.Database

func InitializeDatabase() *mongo.Database {
	once.Do(func() {
		Client = createDbInstance()
	})
	return Client
}

func createDbInstance() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		println("models connection problems")
	}
	var c = client.Database("ums")
	println("models connection establish")
	return c
}

func UserInfoDatabase() *mongo.Collection {
	return Client.Collection("users")
}
