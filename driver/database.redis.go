package driver

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	. "kovercheng/middleware"
)

func newRedisConnection(connectUrl string) error {
	options, err := redis.ParseURL(connectUrl)
	if err != nil {
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "invalid Redis url")
	}
	client := redis.NewClient(options)

	if pong, e := client.Ping(context.Background()).Result(); e != nil || pong != "PONG" {
		Logger.Fatal(e.Error())
		return fmt.Errorf("%s", "cannot ping to Redis database")
	}

	database.Redis = client
	return nil
}

func closeRedisConnection(client *redis.Client) error {
	err := client.Close()
	if err != nil {
		Logger.Fatal(err.Error())
		return fmt.Errorf("%s", "cannot close Redis connection")
	}
	return nil
}
