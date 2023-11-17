package variablecopydemo

import (
	"fmt"
	"testing"
)

func TestVariableReference(t *testing.T) {
	var a, b int = 10, 12
	fmt.Printf("memory of a: %p b: %p\n", &a, &b)
	c := add(a) // a new copy is made and sent to the function
	fmt.Printf("Momory of result of add(a): %p, value: %d\n", &c, c)
}

func TestVariableReferencePtr(t *testing.T) {
	var a, b int = 10, 12
	fmt.Printf("memory of a: %p b: %p\n", &a, &b)
	c := addPtr(&a) // no copy is made the original a's address is passed to the function
	fmt.Printf("Momory of result of add(a): %p, value: %d\n", c, *c)

	addPtr1(&b)
	fmt.Printf("Memory of b: %p ,value of b: %d", &b, b)
}

func add(num int) int {
	fmt.Printf("Memory of argument num: %p\n", &num)
	num++
	fmt.Printf("After increment num++, memory: %p, value: %d\n", &num, num)
	return num //a copy of the num is made and is returned, the local instance of num is destroyed after it is out of scope, (i.e) the function is executed
}

func addPtr(num *int) *int {
	fmt.Printf("Memory of argument num: %p\n", num)
	*num++ //the original value is incremented
	fmt.Printf("After increment num++, momory: %p, value: %d\n", num, *num)
	return num //no copy is made the original value is returned
}

// the value will be changed even if doesn't return
func addPtr1(num *int) {
	fmt.Printf("Memory of argument num: %p\n", num)
	*num++ //the original value is incremented
	fmt.Printf("After increment num++, momory: %p, value: %d\n", num, *num)
	//even if i didn't return the value will be changed
}

// what happens to objects
// the following function demonstrates the at object behaves the same as primitive types
// a copy of the object along with its field is passed to the function , the function again returns the copy of the input
// so totally three copies are made, one local to called function
func TestObjectReference(t *testing.T) {
	var i int = 1
	fmt.Printf("memory of i %p \n", &i) // memory of i 0xc000016388

	c := counter{
		count: &i,
	}
	fmt.Printf("Memory of counter c: %p, c.count: %p\n", &c, c.count)             // Memory of counter c: 0xc00000a048, c.count: 0xc000016388
	d := print(c)                                                                 //a copy of c is made and passsed to the function
	fmt.Printf("d: %p d.count %p, d.count %d, i: %d\n", &d, d.count, *d.count, i) // d: 0xc00000a050 d.count 0xc000016388, d.count 2, i: 2
}

type counter struct {
	count *int
}

func print(obj counter) counter {
	fmt.Printf("Memory of input argument obj: %p, obj.count: %p\n", &obj, obj.count) // Memory of input argument obj: 0xc00000a058, obj.count: 0xc000016388
	*obj.count++
	return obj // a copy of obj is made and returned to the calling function
}

// what hapens to slices when passed to a function as argument
// there are three copies of slice s but the elements themselves are pointers to same object
// so the changes to the elemnts in a function is affected in all slices
func TestSliceReference(t *testing.T) {
	var s []int = make([]int, 5)
	s[0] = 10
	fmt.Printf("mem of s: %p, mem of s[0]: %p \n", &s, &s[0]) // mem of s: 0xc0000080d8, mem of s[0]: 0xc00000e480
	s2 := assign(s)
	fmt.Printf("mem of s2: %p, mem of s2[0]: %p, val of s: %v, val of s2: %v \n", &s2, &s2[0], s, s2) // mem of s2: 0xc0000080f0, mem of s2[0]: 0xc00000e480, val of s: [10 20 0 0 0], val of s2: [10 20 0 0 0]
}

func assign(s1 []int) []int {
	fmt.Printf("mem of s1: %p, mem of s1[0]: %p \n", &s1, &s1[0]) // mem of s1: 0xc000008108, mem of s1[0]: 0xc00000e480
	s1[1] = 20
	return s1
}
