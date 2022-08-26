package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertOne(ctx context.Context, requestId string, request, response []byte, platform int) {
	db := NewMongoDB(ctx)
	coll := db.Collection("logs")
	data := bson.D{
		{Key: "request_id", Value: requestId},
		{Key: "request_body", Value: request},
		{Key: "response_body", Value: response},
		{Key: "created_time", Value: time.Now()},
		{Key: "platform", Value: platform},
	}
	coll.InsertOne(ctx, data)
}
