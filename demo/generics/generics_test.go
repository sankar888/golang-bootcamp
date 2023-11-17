package genericsdemo

import (
	"fmt"
	"testing"
)

func TestGenericsBasics(t *testing.T) {
	ints := []int{1, 2, 3, 4}
	floats := []float32{1.1, 2.1, 3.1, 4.1}
	fmt.Printf("sum of integer: %d\n", sumI(ints))
	fmt.Printf("sum of floats: %v\n", sumF(floats))
	fmt.Printf("sum of ints using generic function: %v\n", sumN(ints))
	fmt.Printf("sum of floats using generic function: %v\n", sumN(floats))

}

// Function without generics
func sumI(ints []int) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return sum
}

// Function without generics
func sumF(floats []float32) float32 {
	var sum float32
	for _, f := range floats {
		sum += f
	}
	return sum
}

// A simple sum function to sum any number, using generics
func sumN[N int | float32](nums []N) N {
	var sum N
	for _, n := range nums {
		sum += n
	}
	return sum
}

/**
* Demo of using generics to get any data, usually a record
 */
func TestGenericData(t *testing.T) {
	queue := make(chan int, 10)
	w := NewWriter(queue)
	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}
		close(queue)
	}()
	w.write()
}

func TestGenericRecord(t *testing.T) {
	queue := make(chan Animal)
	w := NewWriter(queue)
	go func() {
		for i := 0; i < 10; i++ {
			a := Animal{
				species: "Human",
				age:     i,
			}
			queue <- a
		}
		close(queue)
	}()
	w.write()
}

// specific record
type Animal struct {
	species string
	age     int
}

// a generic record
type Record interface {
	Animal | int
}

// a task which takes in a generic record and prints it
type Writer[R Record] struct {
	queue <-chan R
}

func NewWriter[R Record](queue <-chan R) *Writer[R] {
	r := &Writer[R]{
		queue: queue,
	}
	return r
}

func (w *Writer[R]) write() {
	for rec := range w.queue {
		fmt.Println(rec)
	}
}

type Number interface {
	int | int64 | float32 | float64
}

func TestGenericsZerovalue(t *testing.T) {
	print(1)

}

func print[T Number](num T) {
	fmt.Printf("type of number is %T , value: %v \n", num, num)
}
