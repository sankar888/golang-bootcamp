package evaluator

import (
	"errors"
)

func EvaluateExpression(expression string) (res float64, err error) {
	if valid, err := isvalidArithmeticExpression(expression); !valid {
		return 0.0, err
	}
	return 0.0, nil
}

func isvalidArithmeticExpression(expression string) (valid bool, err error) {
	//what is a valid arithmentic expression

	//Rules:
	//operand operator operand. 1 ++++ 2 is not valid
	//the open and close parathesis should match. ((( 2 + 3) * 2) + 3 is not valid
	//( and ) should precede with a operator, unless ( at start and ) at end. 2(2+2) is invalid.
	//there should be a previous ( open before close ). 8 + 2)+( is invalid
	//the start and end should be operand or ( and operand or ). 4 + is invalid
	if length := len(expression); length == 0 {
		return false, errors.New("empty arithmetic expression")
	}
	for i := 0; ; {
		//starts with operand or (

	}
}

// tokens
type token int

const (
	INVALID token = iota
	OPERATORS
	NUMERALS
	OPEN_BRACES
	CLOSE_BRACES
)

type tokenizer struct {
	expression string
}

func newTokenizer(expression string) *tokenizer {
	return &tokenizer{
		expression: expression,
	}
}

func (tzr *tokenizer) nextToken() (t token, val string) {
	
}

type scanner interface {
	read()
}
