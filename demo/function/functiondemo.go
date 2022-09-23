package functiondemo

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

//This will work as the return is changed to float32 type
func mul(x int, y int) float32 {
	return float32(x * y)
}

//return type can have a name
func add(x int, y int) (sum int) {
	return x + y
}

//return type can have a name
func addNakedReturn(x int, y int) (sum int) {
	//sum := x + y //will not work as sum is already declared
	sum = x + y
	return //such return statement is called naked return, the variable declared in return type is returned
}

//the function can have two or more return values
func divide(x int, divideBy int) (int, int) {
	return x / divideBy, x % divideBy
}

func double(x int, y int, z int) (int, int, int) {
	return 2 * x, 2 * y, 2 * z
}

//multiple named return, the argument and return type can't have same name
func swap(x int, y int) (a int, b int) {
	return y, x
}

//a function without argument but return type is allowed
func generateInt() int {
	return 3
}

//if we try to assign the output of this function to a variable or to print it,
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

//this function is not exported so it does not collide with Fn2
func fn2() {
}

//exported function Fn2 is different from fn2
func Fn2() {
}
