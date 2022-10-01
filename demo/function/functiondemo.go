package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

//Entrypoint function - main function in main package should have no argument and return type

//if the return type is not declared and the function has a return go source doesn't compile
/*
func mul(x int, y int) {
	return x * y
}
*/

//go build fails because, the expected return type and actual return type mismatch
/*
func mul(x int, y int) float32 {
	return x * y
}
*/

func main() {
	fmt.Println(add(2, 3))
	mul(3, 5)
	anonymousFunction()
	scopeOfInnerFunctions()
}

// This will work as the return is changed to float32 type
func mul(x int, y int) float32 {
	return float32(x * y)
}

// return type can have a name
func add(x int, y int) (sum int) {
	return x + y
}

// return type can have a name
func addNakedReturn(x int, y int) (sum int) {
	//sum := x + y //will not work as sum is already declared
	sum = x + y
	return //such return statement is called naked return, the variable declared in return type is returned
}

// the function can have two or more return values
func divide(x int, divideBy int) (int, int) {
	return x / divideBy, x % divideBy
}

func double(x int, y int, z int) (int, int, int) {
	return 2 * x, 2 * y, 2 * z
}

// multiple named return, the argument and return type can't have same name
func swap(x int, y int) (a int, b int) {
	return y, x
}

// a function without argument but return type is allowed
func generateInt() int {
	return 3
}

// if we try to assign the output of this function to a variable or to print it,
// it fails: nothing() (no value) used as value
func nothing() {
}

//two functions can't have same name even in a package if the arguments and return type is differnet
/*
func fn1(x int, y int) int {
	return x + y
}

func fn1() {
}
*/

// this function is not exported so it does not collide with Fn2
func fn2() {
}

// exported function Fn2 is different from fn2
func Fn2() {
}

// Anonymous or inner functions
func anonymousFunction() {
	common.Start("Anonymous Function")
	//Anonymous Functions without a name, immediately executed
	func() {
		fmt.Println("iam an anonymous function. Readily executed!")
	}()

	//Anonymous functions can also be assigned to a variable
	a := func() {
		fmt.Println("I am Anonymous!!!, call me to Execute..")
	}
	fmt.Printf("anonymous function type: %T, value: %v \n", a, a) //Type of function is func()

	//Type of function is func()
	var b func() = func() { fmt.Println("hai.. b function called") }
	b()

	//Anonymous functions can have arguments
	c := func(name string) {
		fmt.Println("Hello", name)
	}
	fmt.Printf("c is of type, type : %T \n", c)
	c("Sankar")

	var d func(string) string = func(name string) string {
		return "Hello " + name
	}
	fmt.Printf("calling function d -> d(\"World!\") returns -> %v \n", d("World"))
	common.End()
}

var PI float32 = 2.1432

func scopeOfInnerFunctions() {
	common.Start("InnerFunction Scope")
	var i int = 21
	func() {
		fmt.Println("is i and global scope PI visible inside innder function", i, PI)
		i++
		var j int = 17
		fmt.Println("Value of j delared in inner function", j)
		func() {
			i++
		}()
	}()
	fmt.Println("Value of i after inner function is executed", i)
	//fmt.Println("is Value of j delared in inner function visible outside", j) //No, values declared in inner fuction is only to the scope of that block
	common.End()
}

/**
Learning Resources: https://go101.org/article/function.html
*/
