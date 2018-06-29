package test

import (
	"testing"
	"log"
)

func pfaddPipeline() {
	pipeline := redisClient.Pipeline()

	pipeline.PFAdd("databases", "redis", "mysql", "mongodb")
	pipeline.PFCount("databases")
	_, err := pipeline.Exec()
	if err != nil {
		log.Println("[Error] ", err)
	}
}

func pfadd() {
	cmd := redisClient.PFAdd("databases", "redis", "mysql", "mongodb")
	if err := cmd.Err(); err != nil {
		log.Println("[Error] ", err)
	}
	cmd = redisClient.PFCount("databases")
	if err := cmd.Err(); err != nil {
		log.Println("[Error] ", err)
	}
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
		pfaddPipeline()
	}
}
