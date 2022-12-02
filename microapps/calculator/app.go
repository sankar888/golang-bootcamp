package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/microapps/calculator/evaluator"
)

/*
Calculator app is a cmd line utility which open ups a terminal gets math expressions
*/
func main() {
	printIntroBanner()
	var expression string
	for {
		fmt.Scanln(&expression)
		res, err := evaluator.EvaluateExpression(expression)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(expression, "=", res)
	}
}

var (
	welcomeEmoji string = fmt.Sprintf("%c", '\U0001F64F')
	intro        string = "Welcome to Calculator App. " + welcomeEmoji + newLine +
		`Enter a arithmetic expression to evaluate and hit enter. A valid expression should have only these characters [0-9] [+,-,*,^,%,/,(,),.].`
)

const (
	newLine string = "\n"
)

func printIntroBanner() {
	fmt.Println(intro)
}
