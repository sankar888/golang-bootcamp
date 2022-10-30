package functions

import (
	"fmt"
)

var cache map[int]int = make(map[int]int)

// calculates he fibonacci series of n numbers
func fib(n int) int {
	fmt.Println("Calling fib with inp", n)
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	if n == 0 {
		cache[0] = 0
		return 0
	}
	if n == 1 {
		cache[1] = 1
		return 1
	}
	f := fib(n-1) + fib(n-2)
	cache[n] = f
	fmt.Printf("fib of %d is %d \n", n, f)
	return f
}
