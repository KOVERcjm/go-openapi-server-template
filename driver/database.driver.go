package driver

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	. "kovercheng/middleware"
	"os"
)

type databaseConnection struct {
	Postgres *gorm.DB
	Mongo    *mongo.Client
	Redis    *redis.Client
}

var database databaseConnection

var newConnection = map[string]func(string) error{
	"PGCONNECTURL":    newGormConnection,
	"MONGOCONNECTURL": newMongoConnection,
	"REDISCONNECTURL": newRedisConnection,
}

func init() {
	for db, function := range newConnection {
		if connectUrl, present := os.LookupEnv(db); present {
			if err := function(connectUrl); err != nil {
				Logger.Fatal(err.Error())
				panic(err)
			}
		}
	}
}

func GetConnection() *databaseConnection {
	return &database
}

func CloseConnection() error {
	if database.Postgres != nil {
		if err := closeGormConnection(database.Postgres); err != nil {
			return err
		}
	}
	if database.Mongo != nil {
		if err := closeMongoConnection(database.Mongo); err != nil {
			return err
		}
	}
	if database.Redis != nil {
		if err := closeRedisConnection(database.Redis); err != nil {
			return err
		}
	}
	return nil
}
