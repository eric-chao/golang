package test

import (
	"github.com/go-redis/redis"
	"os"
	"log"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "172.17.68.6:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: 300,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Println("[redis-test] \t", err)
		os.Exit(1)
	}

	log.Println("[redis-test] \t", pong, "connected")

}
