package test

import "testing"

func call() {
	pipeline := redisClient.Pipeline()
	redisClient.Close()
	pipeline.Close()
}

func deferCall() {
	defer redisClient.Close()
	pipeline := redisClient.Pipeline()
	defer pipeline.Close()
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		call()
	}
}

func BenchmarkCallDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deferCall()
	}
}

