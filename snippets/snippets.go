package main

func main() {

}

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

func tokenize(expression string) []token {
	length := len(expression)
	var tokens []token = make([]token, length)
	for i := 0; i < length; i++ {

	}

}
