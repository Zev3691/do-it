package mongo

import (
	"context"
	"re_new/util"
	"re_new/util/constant"
	"re_new/util/log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertOne(ctx context.Context, requestId string, request, response, platform string) {
	db := NewMongoDB().Database("logs")
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

func InsertReqLog(ctx context.Context, request, response string) {
	db := NewMongoDB().Database("logs")
	coll := db.Collection("logs")
	data := bson.D{
		{Key: "request_id", Value: util.GetRequestId(ctx)},
		{Key: "request_body", Value: request},
		{Key: "response_body", Value: response},
		{Key: "created_time", Value: time.Now()},
		{Key: "platform", Value: constant.Req},
	}
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		log.Debug(ctx, err.Error())
	}
}
