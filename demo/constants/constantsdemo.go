package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	constbasicUsage()
	enumsUsingConstants()
}

// create a basic constant
const delimiter byte = ','

//export a constant
const Delimiter byte = delimiter

func constbasicUsage() {
	common.Start("Constant Demo")
	//once declated constant cannot be changed
	const errorCode int = 500
	fmt.Println(errorCode)
	//errorCode = 200 //cannot assign to errorCode (constant 500 of type int)

	//global constants can be shadowed in inner blocks
	const delimiter byte = '|'
	fmt.Printf("value of delimiter inside function %c \n", delimiter)
	func() {
		const delimiter byte = ';'
		fmt.Printf("value of delimiter inside anonynows function %c \n", delimiter)
	}()

	//untyped constants
	const a = 1
	fmt.Printf("type of a is infered %v, %T \n", a, a)
	var b int8 = 10
	c := b + a //infered constants can be used with similar types (int8 + inferred constant)
	fmt.Printf("inferred constant addition %v, %T \n", c, c)
	common.End()
}

type color int8

const (
	_ color = iota
	red
	green
	blue
)

type method string

const (
	get    method = "GET"
	post   method = "POST"
	put    method = "PUT"
	delete method = "DELETE"
)

func enumsUsingConstants() {
	common.Start("Enums")

	//constants enumreated using iota can be used as enums
	//var c int8 = 1
	//fmt.Printf("color %v is equalt to blue : %v \n", c, c == blue) //invalid operation: c == blue (mismatched types int8 and color)
	var c1 color = red
	fmt.Printf("color red is equal to blue : %v \n", c1 == blue)

	var m method = put
	fmt.Printf("method m : %v", m)
	common.End()
}
