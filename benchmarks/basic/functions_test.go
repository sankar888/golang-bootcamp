package functions

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fib(10)
}

func BenchmarkFib(b *testing.B) {
	fmt.Println(b.N)
}
