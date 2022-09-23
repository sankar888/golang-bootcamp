package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	basics()
	basics0()
	chaining()
	var a, b int = 10, 20
	fmt.Println(swap(&a, &b))
	common.End()
	pointerValueChange()
}

func basics() {
	common.Start("basics")
	// A pointer holds the memory address of a value
	var i int = 10
	p := &i //& operator get the addres of its operand, &i get the address of the variable i
	fmt.Printf("The variable p is a pointer. It holds an address %v \n", p)
	fmt.Printf("The type of pointer p is %T \n", p)
	fmt.Printf("The value pointed by the pointer is %v \n", *p) //* operator de-refernce the pointer and gets the value  pointed by the pointer
	common.End()
}

func basics0() {
	common.Start("basics0")
	//A pointer can be delcared as follows
	//var p *int = 0xc000016088 //this throws build error because we cannot assign memory address
	var p *int
	fmt.Printf("the unassigned pointer p has address %v \n", p) //unassigned pointer has address nil
	//fmt.Printf("the value of unassigned pointer is %v \n", *p) //panic: runtime error: invalid memory address or nil pointer dereference
	var ptr *int8 = nil
	fmt.Printf("the unassigned pointer ptr has address %v \n", ptr) //unassigned pointer has address nil
	common.End()
}

// chaining pointers
func chaining() {
	common.Start("chaining")
	var i int8 = 15
	var p *int8 = &i
	var pp **int8 = &p
	fmt.Printf("chaining pointers is valid in go: %v \n", **pp) //this is valid but dont use it as it not practical
	fmt.Printf("chaining pointers is valid in go: %v \n", *p)
	common.End()
}

// pointers can be asigned another address
func swap(x *int, y *int) (int, int) {
	common.Start("swap")
	var tmp *int = x
	x = y
	y = tmp
	return *x, *y
}

// thevalue of the pointers can be changed or manipulated
func pointerValueChange() {
	common.Start("pointervalueChange")
	var i int8 = 22
	p := &i
	*p = *p + 1
	fmt.Printf("the initial value of i is 22, after change using pointers: %v \n", i)
	fmt.Printf("the initial value of i is 22, After change, accessing using poitnters %v \n", *p)
	i = 19
	fmt.Printf("The pointer points to %v \n", *p)
	common.End()
}
