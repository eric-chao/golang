package redis

import (
	. "adhoc/adhoc_data_fast/config"
	. "adhoc/adhoc_data_fast/logger"
	"github.com/go-redis/redis"
	"fmt"
)

var client redis.Client

func init() {
	redisAddress := fmt.Sprintf("%s:%s", GlobalConfig.Redis.Host, GlobalConfig.Redis.Port)
	&client = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: GlobalConfig.Redis.PoolSize,
	})
}

func GetKey(key string) (string, error) {
	val, err := client.Get(key).Result()
	if err != nil {
		Logger.Error("[Redis]: get key error!")
		return "", err
	}
	return val, nil
}
