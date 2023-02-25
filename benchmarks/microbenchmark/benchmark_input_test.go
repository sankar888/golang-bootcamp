package microbenchmark

import (
	"testing"
)

/*
size of input 1
size of input 100
size of input 10000 - 100k
size of input 1000000 - 1M
size of input 100000000 - 100M
size of input 1000000000 - 1B
*/
func BenchmarkInput(b *testing.B) {
	b.Logf("size of input %d\n", b.N)
}
