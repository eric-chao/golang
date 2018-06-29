package test

import (
	"testing"
	"bytes"
)

func hincWithPipeline(i int) {
	pipeline := redisClient.Pipeline()
	//defer pipeline.Close()

	var buf bytes.Buffer
	buf.WriteString("r")
	buf.WriteString(string(i))

	pipeline.HIncrBy("key", buf.String(), 1)
	pipeline.Exec()

}

func hinc(i int) {

	var buf bytes.Buffer
	buf.WriteString("r")
	buf.WriteString(string(i))

	redisClient.HIncrBy("key1", "redis", 1)

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
