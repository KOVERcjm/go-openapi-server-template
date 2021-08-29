package database

import (
	"context"
	"kovercheng/driver"
	. "kovercheng/middleware"
)

var redis = driver.GetConnection().Redis

func redisErrorHandler(err error) {
	Logger.Infof("Redis query ERROR: %s", err)
}

func TestRedis() error {
	if err := redis.Set(context.Background(), "key", 123, 0).Err(); err != nil {
		redisErrorHandler(err)
	}
	Logger.Infoln("Redis insert success.")

	if err := redis.Set(context.Background(), "key", 456, 0).Err(); err != nil {
		redisErrorHandler(err)
	}
	Logger.Infoln("Redis update success.")

	findResult, e := redis.Get(context.Background(), "key").Result()
	if e != nil {
		redisErrorHandler(e)
	}
	Logger.Infof("Redis find success: %+v", findResult)

	if err := redis.Del(context.Background(), "key").Err(); err != nil {
		redisErrorHandler(err)
	}
	Logger.Infoln("Redis delete success.")

	return nil
}
