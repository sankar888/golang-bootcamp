package breaksandcontinue

import (
	"testing"
)

// this function demonstrates the basic usage of break
// In Golang, a break statement can be used for terminating execution of innermost for, switch or select statement, 
// after that it transfers the execution of program to the next statements following the break statement.
// Refer: https://yourbasic.org/golang/switch-statement/
// https://go.dev/ref/spec#Break_statements
func TestBreakBasicUsage(t *testing.T) {
	// break is valid only in a loop, switch or select context
	if 1 == 1 {
		//break // will fail with: "break is not in a loop, switch, or select"
	}
	//break //will fail

	//break in loop. break will break only the innermost loop
	times := 0
	for i := 1; i <= 10; i++ { //will run 10 times
		for j := 1; j <= 10; j++ { //for each time it will run i times, ie 1 + 2 + 3 + .. 10 = 10 * 11 / 2 = 55 times
			t.Logf("i, j -> %d, %d\n", i, j)
			times += 1
			if i == j {
				t.Log("breaking innermost loop.")
				break
			}
		}
	}
	if times != 55 {
		t.Logf("Expected to run %d times, actual loop executions %d \n", 55, times)
		t.Fail()
	}

	// switch statements evaluates every case one by one from top to bottom and executes the first matching case. 
	// the rest of the cases are not evaluated. So why is break statement needed in switch
	// break statement is not needed for one line switch cases. it is needed to break from the logic as shown below
	swfn := func(i int) {
		switch {
		case i <= 3:
			if i == 2 {
				t.Log("i is 2. breaking. the statements below will not execute.")
				break
			}
			t.Logf("%d is <= 3\n", i) // will not be printed if i is 2
		case i > 3:
			t.Logf("%d is > 3\n", i)	
		}
	}
	swfn(3)
	swfn(2)
}

// this function demonstrates the usage of break statement with labels.
// go has goto and labels, which passes the control flow of the program
// A labeled statement may be the target of a goto, break or continue statement.
// Labels are declared by labeled statements and are used in the "break", "continue", and "goto" statements. 
// It is illegal to define a label that is never used. In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. 
// The scope of a label is the body of the function in which it is declared and excludes the body of any nested function
// Refer: https://go.dev/ref/spec#Break_statements
func TestBreakWithlabels(t *testing.T) {
	//start: //gives invalid break label start as the label which is referred by break and continue starements should refer a loop, select or switch statement
	times := 1
	outer:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			t.Logf("i, j -> %d, %d\n", i, j)
			times += 1
			if i == j {
				t.Log("breaking to outermost loop.")
				break outer
			}
		}
	}
	t.Log("broke the loop. exited")
}

//Add break in select