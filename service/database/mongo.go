package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"kovercheng/driver"
	. "kovercheng/middleware"
	"kovercheng/model/document"
)

var mongo = driver.GetConnection().Mongo.Database("exampleMongo")

func mongoErrorHandler(err error) {
	Logger.Infof("Mongo query ERROR: %s", err)
}

func TestMongo() error {
	testCollection := mongo.Collection("test")

	if _, err := testCollection.InsertOne(context.TODO(), document.Test{Key: "key", Value: 123}); err != nil {
		mongoErrorHandler(err)
		return fmt.Errorf("mongo query error")
	}
	Logger.Infoln("Mongo insert success.")

	if _, err := testCollection.UpdateOne(context.TODO(), bson.M{"key": "key"}, bson.M{"$set": bson.M{"value": 456}}); err != nil {
		mongoErrorHandler(err)
		return fmt.Errorf("mongo query error")
	}
	Logger.Infoln("Mongo update success.")

	var findResult document.Test
	result := testCollection.FindOne(context.TODO(), bson.M{"key": "key"})
	if result.Err() != nil {
		mongoErrorHandler(result.Err())
		return fmt.Errorf("mongo query error")
	}
	_ = result.Decode(&findResult)
	Logger.Infof("Mongo find success: %+v", findResult)

	if _, err := testCollection.DeleteOne(context.TODO(), bson.M{"key": "key"}); err != nil {
		mongoErrorHandler(err)
		return fmt.Errorf("mongo query error")
	}
	Logger.Infoln("Mongo delete success.\n")

	return nil
}
