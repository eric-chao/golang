package redis

import (
	. "adhoc/adhoc_data_fast_golang/config"
	"github.com/go-redis/redis"
	"fmt"
)

var DataRedisClient *redis.Client

func init() {
	redisAddress := fmt.Sprintf("%s:%s", GlobalConfig.Redis.DataHost, GlobalConfig.Redis.DataPort)
	DataRedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: GlobalConfig.Redis.ExpPoolSize,
	})
}
