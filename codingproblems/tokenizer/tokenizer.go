package main

func main() {

}

/*The aim is to tokenize a simple arithmetic expression
  The arithmetic expression can have
  - operand (decimal and whole numbers)
  - operators (+, -, *, /)
*/

type tokenType int8

const (
	operand tokenType = iota + 1
	operator
)

type token struct {
	data  float64
	ttype tokenType
}

func (t *token) isOperand() bool {
	return t.ttype == operand
}

func (t *token) isOperator() bool {
	return t.ttype == operator
}

func (t *token) getOperand() float64 {
	return t.data
}

func (t *token) getOperator() byte {
	return byte(t.data)
}

type state int8

// State
const (
	makeOperand state = iota + 1
	makeOperator
	tokenize
)

type queue struct {
	//TODO: create queue datastructure
}

func tokenizeExp(expression string) []token {
	//length := len(expression)
	//var tokens []token = make([]token, length)

	return nil
}
