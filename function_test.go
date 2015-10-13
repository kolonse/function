package function

import (
	"testing"
)

func test1(t *testing.T, hh string) {

}

func TestCall(t *testing.T) {
	f := Bind(test1, P_1, "haha")
	f.Call(t)
}

func test2(t *testing.T, hh string) {
}

func TestCall2(t *testing.T) {
	f := Bind(test1, P_1, P_2)
	f.Call(t, "test2")

}

func bench(b *testing.B, hh string) {
}

func BenchmarkSpeed1(b *testing.B) {
	f := Bind(bench, P_1, P_2)
	for i := 0; i < b.N; i++ {
		f.Call(b, "test2")
	}
}

func BenchmarkSpeed2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bench(b, "test2")
	}
}
