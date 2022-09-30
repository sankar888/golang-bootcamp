package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	declaringVariables()
}

// declare and initiate variables
func declaringVariables() {
	common.Start("**variables Basics Demo**")
	var i int = 12 //var [variable_identifier] [variable_type] = [value]
	fmt.Printf("var i int 12, value of i is %v \n", i)
	//var i int = 4 //A varible can be declared only one time in a specific scope

	var j int
	fmt.Printf("var j int, value of uninitialized j is %v \n", j)
	j = 4
	fmt.Printf("j = 4, once a variable is declared its value can be changed, j : %v \n", j)

	//short hand notation
	k := 23.4
	fmt.Printf("k := 23.4, variable can be also declared with short hand notation, the type is infered, k : %v, type of k: %T \n", k, k)

	var a, b, c int //multiple variables can be delcared at one
	fmt.Printf("var a, b, c, value of a, b, c is %v, %v, %v \n", a, b, c)
	a, b, c = 1, 2, 3
	fmt.Printf("a, b, c = 1, 2, 3 multiple variable can also be initialized at once, value of a, b, c : %v, %v, %v \n", a, b, c)

	x, y := 12, 13.5
	fmt.Printf("x, y := 12, 13.5 multiple variables can be initialized with short hand notations, value of x, y = %v, %v, type of x, y = %T, %T\n", x, y, x, y)
	y = 12
	fmt.Printf("Value of y can be changed, but type once initilized cant be changed, %v, %T \n", y, y)
	common.End()
}
