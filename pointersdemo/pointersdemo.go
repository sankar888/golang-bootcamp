package main

import (
	"fmt"
)

func main() {
	basics()
	basics0()
	chaining()
}

func basics() {
	// A pointer holds the memory address of a value
	var i int = 10
	p := &i //& operator get the addres of its operand, &i get the address of the variable i
	fmt.Printf("The variable p is a pointer. It holds an address %v \n", p)
	fmt.Printf("The type of pointer p is %T \n", p)
	fmt.Printf("The value pointed by the pointer is %v \n", *p) //* operator de-refernce the pointer and gets the value  pointed by the pointer
	fmt.Println("----")
}

func basics0() {
	//A pointer can be delcared as follows
	//var p *int = 0xc000016088 //this throws build error because we cannot assign memory address
	var p *int
	fmt.Printf("the unassigned pointer p has address %v \n", p) //unassigned pointer has address nil
	//fmt.Printf("the value of unassigned pointer is %v \n", *p) //panic: runtime error: invalid memory address or nil pointer dereference
	var ptr *int8 = nil
	fmt.Printf("the unassigned pointer ptr has address %v \n", ptr) //unassigned pointer has address nil
	fmt.Println("----")
}

// chaining pointers
func chaining() {
	var i int8 = 15
	var p *int8 = &i
	var pp **int8 = &p
	fmt.Printf("chaining pointers is valid in go: %v \n", **pp) //this is valid but dont use it as it not practical
	fmt.Printf("chaining pointers is valid in go: %v \n", *p)
}
