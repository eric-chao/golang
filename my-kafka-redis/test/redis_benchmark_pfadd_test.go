package test

import (
	"testing"
)

func pfaddPipeline() {
	pipeline := redisClient.Pipeline()
	//defer pipeline.Close()

	pipeline.PFAdd("databases", "redis", "mysql", "mongodb")
	pipeline.PFCount("databases")
	pipeline.Exec()
	redisClient.Close()
}

func pfaddPipelineDefer() {
	defer redisClient.Close()
	pipeline := redisClient.Pipeline()
	//defer pipeline.Close()

	pipeline.PFAdd("databases", "redis", "mysql", "mongodb")
	pipeline.PFCount("databases")
	pipeline.Exec()
}

func pfadd() {
	redisClient.PFAdd("databases", "redis", "mysql", "mongodb")
	redisClient.PFCount("databases")
	redisClient.Close()
}

func pfaddDefer() {
	defer redisClient.Close()

	redisClient.PFAdd("databases", "redis", "mysql", "mongodb")
	redisClient.PFCount("databases")

}

func Benchmark_pfadd(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		pfadd()
	}
}

func Benchmark_pfadd_defer(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		pfaddDefer()
	}
}

func Benchmark_pfadd_pipeline(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		pfaddPipeline()
	}
}

func Benchmark_pfadd_pipeline_defer(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		pfaddPipelineDefer()
	}
}
