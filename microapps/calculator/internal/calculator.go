// Calculator evaluates the given expression and returns the result
// It also stores the last x successful calculated expressions and its results, x is configured when a calculator type is created
package calculator

import (
	"errors"
	"strings"
)

type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

// Eval evaluates the input arithmetic expressions
// return a single number or error if the expression cannot be evaluated
func Eval(expression string) (float64, error) {
	trimmedExpression := strings.TrimSpace(expression)
	length := len(trimmedExpression)
	if length == 0 {
		return 0, errors.New("empty arithmetic expression")
	}

	return 0, nil
}

func infix2posfix(expression string) string {
	length := len(expression)
	for i := 0; i < length; i++ {

	}
	return ""
}

func evaluatePosfixExpression(pexpression []byte) float64 {
	return 0.0
}

const (
	Numeric  int8 = 0
	Operator int8 = 1
)

type token struct {
	startIndex int8
	endIndex   int8 //exclusive
	tokenType  int8 //constant
}

func tokenize(expression string) []token {
	//what would be the token type be
	//token could be a struct or can the operators are actually characters which can be held in float64
	length := len(expression)
	var tokens []token = make([]token, length)

}
