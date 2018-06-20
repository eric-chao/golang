package test

import (
	"testing"
	"bytes"
)

func hincWithPipeline(i int) {
	redis := NewRedisClient()
	defer redis.Close()
	pipeline := redis.Pipeline()
	defer pipeline.Close()

	var buf bytes.Buffer
	buf.WriteString("r")
	buf.WriteString(string(i))

	pipeline.HIncrBy("key", buf.String(), 1)
	pipeline.Exec()

}

func hinc(i int) {
	redis := NewRedisClient()
	defer redis.Close()

	var buf bytes.Buffer
	buf.WriteString("r")
	buf.WriteString(string(i))

	redis.HIncrBy("key1", "redis", 1)

}

func Benchmark_hinc(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		hinc(i)
	}
}

func Benchmark_hincWithPipeline(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		hincWithPipeline(i)
	}
}
