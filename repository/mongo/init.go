package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoDb *mongo.Client

func Init() {
	mongoDb = initMongo()
}

func initMongo() *mongo.Client {
	uri := "mongodb://admin:123456@localhost:27017/?maxPoolSize=20&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	println("[mongo] Successfully connected and pinged.")

	return client
}

func NewMongoDB() *mongo.Client {
	return mongoDb
}
