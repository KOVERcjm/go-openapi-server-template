package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	. "kovercheng/middleware"
	"time"
)

func newMongoConnection(connectUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(connectUrl))
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		_ = closeMongoConnection(client)
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "cannot ping to database")
	}
	database.Mongo = client
	return nil
}

func closeMongoConnection(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "cannot close Mongo connection")
	}
	return nil
}
