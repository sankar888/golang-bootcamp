package random

import (
	"math/rand"
	"testing"
)

func TestRandomInGo(t *testing.T) {
	var min, max int = 10, 30
	for i := 0; i < 100; i++ {
		t.Log(rand.Intn(int(max-min)) + min)
	}
}

func TestRandomFloats(t *testing.T) {
	var min, max float64 = -5.50, 5.50
	for i := 0; i < 100; i++ {
		t.Log((rand.Float64() * (max - min)) + min)
	}
}
