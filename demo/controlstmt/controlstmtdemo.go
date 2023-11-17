package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	basicIf()
	basicIfwithinit()
	ifElse()
	ifElseChain()
	basicSwitch()
	switchWithNoCondition()
	switchExecutionOrder()
	switchWithMultipleCase()
	switchWithMultipleCorrectCase()
	switchWithFallthrough()
}

// basic if
func basicIf() {
	common.Start("**Basic If Demo**")
	var i int = 101
	if i > 100 { //braces are mandatory
		fmt.Println("Big Number")
	}
	common.End()
}

func basicIfwithinit() {
	common.Start("**IF With init Block**")
	if i := 101; i > 100 { //like loops, if also can have a init block, the scope of init block is within the if.. else statement
		fmt.Println("Big Number")
	}
	//fmt.Println(i) //demo\controlstmt\controlstmtdemo.go:29:14: undefined: i
	common.End()
}

func ifElse() {
	common.Start("**If..Else Demo**")
	if i := 7; i%2 == 0 { //the init block is visible in the whole if..else chain
		fmt.Printf("%v is an even number \n", i)
	} else {
		fmt.Printf("%v is an odd Number \n", i)
	}
	common.End()
}

func ifElseChain() {
	common.Start("**If..Else Chain Demo**")
	if os := runtime.GOOS; os == "ubuntu" {
		fmt.Println("This mahine is running UBUNTU os")
	} else if os == "linux" {
		fmt.Println("This mahine is running linux os")
	} else if os == "windows" {
		fmt.Println("This mahine is running windows os")
	} else {
		fmt.Println("This mahine is running unknown os")
	}
	common.End()
}

// switch
func basicSwitch() {
	//switch canhandle any type, not only constants
	//switch also can have a init block
	//the scope of init block is with the switch statement
	common.Start("**Switch stmt Demo**")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("This is a linux machine")
		fmt.Println("This is a linux machine") //no need for break statement
	case "windows":
		{
			fmt.Println("This is a windows machine") //the case is executed in the order defined one after another
			fmt.Println("case block can have multiple statement")
		}
	default:
		fmt.Println("unknown os")
	}
	//fmt.Println(os) //will throw undefined error
	common.End()
}

func switchWithNoCondition() {
	common.Start("**Switch with NoCondition Demo**")
	i := 72
	switch {
	case i >= 1 && i <= 10:
		fmt.Printf("%v is in Range 1 .. 10 \n", i)
	case i >= 10 && i < 20:
		fmt.Printf("%v is in Range 10 .. 20 \n", i)
	case i >= 30 && i < 50:
		fmt.Printf("%v is in Range 30 .. 50 \n", i)
	default:
		fmt.Printf("%v is in Range > 50 \n", i)
	}
	common.End()
}

func switchExecutionOrder() {
	common.Start("**Switch with NoCondition Demo**")
	i := 12 //initial value of i
	switch {
	case change(&i, 31) && i >= 1 && i <= 10: //All switch cases are evaluated in the order they appear, All three cases are evaluated before going to default
		fmt.Printf("%v is in Range 1 .. 10 \n", i)
	case i >= 10 && i < 20:
		fmt.Printf("%v is in Range 10 .. 20 \n", i)
	case i >= 30 && i < 50:
		fmt.Printf("%v is in Range 30 .. 50 \n", i)
	default:
		fmt.Printf("%v is in Range > 50 \n", i)
	}
	common.End()
}

func change(p *int, newval int) bool {
	*p = newval
	return false
}

func switchWithMultipleCase() {
	common.Start("Switch with same logic for multiple cases")
	switch today := time.Now().Weekday(); today {
	case time.Sunday, time.Saturday: //multiple case points to same logic
		fmt.Printf("today is %s. No work!\n", today)
	default:
		fmt.Printf("today is %s, Have to work\n", today)
	}
	common.End()
}

// switch will evaluate the cases in the order they appear, from top to bottom and executes the first matched case. other cases are not evaluated
func switchWithMultipleCorrectCase() {
	common.Start("Multiple correct switch case")
	switch bugs := 10; {
	case bugs > 3:
		fmt.Printf("bugs %d is > 3\n", bugs)
	case fact() && bugs > 5: //since the first case is true the remaining cases are not evaluated
		fmt.Printf("bugs %d is > 5\n", bugs)
	}
	common.End()
}

func fact() bool {
	fmt.Println("this function will always return true")
	return true
}

// switch will evaluate the cases in the order they appear, from top to bottom and executes the first matched case. other cases are not evaluated
// but we have special fallthrough statment which passes the control to next case
// fallthrough cannot be used with a type switch
// fallthrough will execute the body of the next case, no checking that next case for a match!
func switchWithFallthrough() {
	common.Start("Switch with pass through")
	switch bugs := 10; {
	case bugs > 3:
		fmt.Printf("bugs %d is > 3\n", bugs)
		fallthrough //fallthrough passes the control to the body of the next case without evaluating the case expression
		//fallthrough should be the last line of the case body
	case fact() && bugs > 5: //this case expression will not be evaluated, but the case boday will be executed.
		fmt.Printf("bugs %d is > 5\n", bugs)
	default:
		fmt.Println("default case.")
	}
	common.End()
}
