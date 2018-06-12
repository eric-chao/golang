package redis

import (
	. "adhoc/adhoc_data_fast/config"
	"github.com/go-redis/redis"
	"fmt"
)

func NewDataRedisClient() *redis.Client {
	redisAddress := fmt.Sprintf("%s:%s", GlobalConfig.Redis.ExpHost, GlobalConfig.Redis.ExpPort)
	DataRedisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: GlobalConfig.Redis.ExpPoolSize,
	})
	return DataRedisClient
}
