package redis

import (
	. "golang/my-kafka-redis/config"
	"github.com/go-redis/redis"
	"fmt"
	"os"
	"log"
)

var dataRedis *redis.Client

func init() {
	redisAddress := fmt.Sprintf("%s:%s", GlobalConfig.Redis.DataHost, GlobalConfig.Redis.DataPort)
	dataRedis = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: GlobalConfig.Redis.ExpPoolSize,
	})

	pong, err := dataRedis.Ping().Result()
	if err != nil {
		log.Println("[redis-data] \t", err)
		os.Exit(1)
	}

	log.Println("[redis-data] \t", pong, "connected")
}

func GetDataRedis() *redis.Client {
	return dataRedis
}

func CloseDataRedis() {
	dataRedis.Close()
}

