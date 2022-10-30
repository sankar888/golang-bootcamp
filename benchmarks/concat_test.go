package concat

import (
	"testing"
)

const s string = "History"

func BenchmarkStringAddition(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = res + s
	}
	res = ""
}

func BenchmarkStringConcatBuffer(b *testing.B) {
	var res []byte = make([]byte, 0)
	for i := 0; i < b.N; i++ {
		res = append(res, s...)
	}
	res = nil
}
