package test

import (
	"testing"
	"github.com/go-redis/redis"
	"strings"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.216.175:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: 300,
	})
	return client
}

func pfaddWithPipeline() {
	redis := NewRedisClient()
	defer redis.Close()
	pipeline := redis.Pipeline()
	defer pipeline.Close()

	pipeline.PFAdd("databases", "redis", "mysql", "mongodb")
	pipeline.PFCount("databases")
	cmd, err := pipeline.Exec()
	if err == nil {
		strings.Split(cmd[1].String(), " ")
	}
}

func pfadd() {
	redis := NewRedisClient()
	defer redis.Close()

	redis.PFAdd("databases", "redis", "mysql", "mongodb")
	redis.PFCount("databases")

}

func Benchmark_pfadd(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		pfadd()
	}
}

func Benchmark_pfadd_pipeline(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		pfaddWithPipeline()
	}
}
