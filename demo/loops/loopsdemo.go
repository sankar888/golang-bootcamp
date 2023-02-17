package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	basicLoop()
	loopOptionals()
}

func basicLoop() {
	common.Start("**Basic Loop Demo**")
	//go has only one looping constraint. for
	//for has three separate parts separated by semicolons
	//1. the initializer - optional
	//2. the condition
	//3. the increment statement - optional
	var sum int
	for i := 0; i < 10; i++ { //braces are mandatory
		//long declaration like var i int = 0 is not allowed in for loop initialization
		//this is by design to make initializations simple
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	for i, until := 1, 10; i < until; i++ {
		//i and until exists in block scope, it is not accessible outside this loop
		//initializtion block is run only once
		sum += i
	}
	fmt.Println(sum)
	//fmt.Println(i) //demo\loops\loopsdemo.go:35:14: undefined: i
	common.End()
}

func loopOptionals() {
	common.Start("**Loop Optionals Demo**")
	//intializtion and increment is optional in go loop
	var count int8 = 10
	for ; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Printf("loopng demo - loop optionals - %v \n", count)

	count = 10
	for count > 0 { //if there is not incremental or initialization statement, then the semicolons can be left out and for can be used like while loop
		count--
	}
	fmt.Printf("loopng demo - for as while - %v \n", count)
	common.End()
}

// infinite loop
func forever() {
	//a loop which has no condition is always true and runs infinitely
	for {

	}
}
