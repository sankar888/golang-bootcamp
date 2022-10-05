package main

func main() {

}

/*The aim is to tokenize a simple arithmetic expression
  The arithmetic expression can have
  - operand (decimal and whole numbers)
  - operators (+, -, *, /)
*/

const (
	operand int8 = iota
	operator
)

type token struct {
	data  float64
	ttype int8
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

// State
const (
	completeOperand int8 = iota
	//TODO: Effective use of constants and enums in golang
)

func tokenize(expression string) []token {
	return nil
}
