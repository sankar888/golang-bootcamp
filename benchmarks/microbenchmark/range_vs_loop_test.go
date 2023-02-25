package microbenchmark

import (
	"testing"
)

/*
*
The below benchmarks compare what is faster
1. using range to loop through array and slices
2. or using for loop and index

Observation:

1. For constant 100k num sized array, 20s benchtime

There is neglible difference between the two for 100k size array, The index based approach has a vey minute 30 ns difference per operation.
BenchmarkRangeBasedIteration-8           8700958              2571 ns/op               0 B/op          0 allocs/op
BenchmarkIndexBasedIteration-8           8963062              2531 ns/op               0 B/op          0 allocs/op

2. For 1M array size, benchtime 20s
negligible difference  1 ms, but this time in favor of range based iteration
BenchmarkRangeBasedIteration-8            675177             29866 ns/op               2 B/op          0 allocs/op
BenchmarkIndexBasedIteration-8            695733             30838 ns/op               2 B/op          0 allocs/op

A good read regarding for loop and range based iterations : https://medium.com/@ddwen/handling-large-arrays-in-golang-should-you-use-for-range-or-for-loop-9995a02fd316
It says that when looping over large array or slices of structs which has more than 4 fields the for-index loops will be faster than rage, coz range makes copies of object
*/
const (
	size int = 10 * 100 * 100
)

func BenchmarkRangeBasedIteration(b *testing.B) {
	//size := b.N
	nums := createNumericSlice(size)
	var out []int = make([]int, size)
	b.Log(size)
	b.Logf("%d times executed array size %d\n", b.N, size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i, num := range nums {
			out[i] = num
		}
	}
	b.Log(out[0])
	b.StopTimer()
}

func BenchmarkIndexBasedIteration(b *testing.B) {
	//size := b.N
	nums := createNumericSlice(size)
	var out []int = make([]int, size)
	b.Logf("%d times executed array size %d\n", b.N, size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < size; i++ {
			out[i] = nums[i]
		}
	}
	b.Log(out[0])
	b.StopTimer()
}

func createNumericSlice(size int) []int {
	if size <= 0 {
		panic("cannot create a slice of length <= 0")
	}
	var nums []int = make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}
	return nums
}

func TestRangeMakeCopiesOfObjet(t *testing.T) {
	type ticket struct {
		id string
	}

	//the array of tickets is not affected by the assignment inside the range loop, coz the tkt object provided by range is a copy of original object in tickets array
	var tickets []ticket = make([]ticket, 10)
	size := len(tickets)
	t.Log(size)
	for i, tkt := range tickets {
		if i == 3 {
			tkt.id = "tkt_id"
		}
	}
	t.Logf("array of tickets %v\n", tickets)

	//index based loop affects the original object in array, bcoz it modifies object directly
	for i := 0; i < size; i++ {
		if i == 3 {
			tickets[i].id = "ticket_id_3"
		}
	}
	t.Logf("array of tickets after for loop %v\n", tickets)

	//conclusion: index based for loop can be used for both writing and reading from slices, maps, channels etc..
	//range based loops can be used for reading and it makes  copy
}
