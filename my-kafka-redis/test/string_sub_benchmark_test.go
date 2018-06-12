package test

import (
	"testing"
)

func sub_string1() {
	s := "ADHOC_d1a5fa92-d874-48d6-bc9b-46056563e671"
	s = string([]byte(s)[6:])
}

func sub_string2() {
	s := "ADHOC_d1a5fa92-d874-48d6-bc9b-46056563e671"
	s = string([]rune(s)[6:])
}

func Benchmark_sub_string1(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		sub_string1()
	}
}

func Benchmark_sub_string2(b *testing.B) {
	//use b.N for looping
	for i := 0; i < b.N; i++ {
		sub_string2()
	}
}