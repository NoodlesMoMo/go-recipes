package parse

import (
	"testing"
)

func Benchmark_Unpack(b *testing.B) {
	b.StopTimer()
	u := "name=xiaoming&score=123"

	obj := &URLParam{}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ParamsUnpack(obj, u)
	}
}

func Benchmark_Unpack2(b *testing.B) {
	b.StopTimer()
	u := "name=xiaoming&score=123"

	obj := &URLParam{}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ParamsUnpack2(obj, u)
	}
}
