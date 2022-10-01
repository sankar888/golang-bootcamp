package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	arraysIntro()
	arraysContd()
	multiDimensionalArray()
	arrayIsOfvalueType()
}

// Array is a fixed length immutable homogenous data structure in Go language
func arraysIntro() {
	common.Start("**Arrays in Go**")
	var arr [5]int = [5]int{} // This is how array is declared, an uninitialized array has nil values for the type and fixed length
	fmt.Printf("array 'arr' initial values %v, arraylength %d \n", arr, len(arr))

	arr0 := [4]float32{} //could be declared with shorthand notation
	fmt.Printf("'arr0' is of type %T, value: %v \n", arr0[0], arr0)

	//arr = [3]int{1, 2, 3} // This is not possible, because in go array length is also part of type, type(arroflen3) != type(arroflen4)

	arr = [5]int{1, 2, 4, 8, 16}
	fmt.Printf("array 'arr' values %v, arraylength %d \n", arr, len(arr))

	//Arrays values can be changed but array size cannot be extended or changed
	arr[0] = 101
	fmt.Printf("array 'arr' value changed %v \n", arr)
	//fmt.Println(arr[7]) //invalid argument: index 7 out of bounds [0:5]

	//Length of an array
	fmt.Printf("len(arr) gives the lenght of array %d \n", len(arr))
	common.End()
}

func arraysContd() {
	common.Start("**Arrays Continued..**")
	arr := [...]int{} //we could also omit the length of a array and ask the compiler to compute the same
	fmt.Printf("'arr' value: %v, length: %d \n ", arr, len(arr))

	arr1 := [...]int{1, 2, 3}
	fmt.Printf("'arr1' value: %v, length: %d \n ", arr1, len(arr1))

	//arr = [3]int{1, 2, 3} //not possible , coz arr is of type [0]int cannot assign [3]int
	common.End()
}

func multiDimensionalArray() {
	common.Start("MultiDimensional Array")
	var board [4][5]int8 = [4][5]int8{} //i, j => i denotes row, j denotes column
	fmt.Printf("2 dimensional array %v \n", board)

	//how to get the dimensions of multidimensional array
	fmt.Printf("length of multidimensional array, columnlength (i): %d, rowlength(j) : %d \n", len(board), len(board[0]))

	//can multidimensional array have rows of different length ? No..
	var board1 [2][2]int8 = [2][2]int8{{1, 2}, {2, 3}}
	fmt.Printf("Multidimensional array can be initialized with literals value: %v \n", board1)
	board1[1] = [2]int8{0, 0}
	fmt.Printf("values can be changed. values: %v \n", board1)

	//Multi dimensional array can use ... notation for size
	//board2 := [...][...]int{{1}, {2}, {3}} //not working

	//How to access multidimensional array values
	fmt.Printf("values of multidimensional array canbe accessed as follows, arr[i][j]: %v \n", board1[0][1])

	//More than two dimensional array is possible
	var cube [2][3][4]int8 = [2][3][4]int8{} //i, j, k -> 2, 3, 4
	fmt.Printf("three dimensioanl array, value: %v \n", cube)
	common.End()
}

// Array in go is of value type, when it is passed to another function as argument, a copy is send
func arrayIsOfvalueType() {
}
