package breaksandcontinue

import (
	"testing"
	"github.com/sankar888/golang-bootcamp/demo/common"
)

// Refer: https://go.dev/ref/spec#Continue_statements
// Unlike break, a continue statement is valid only in loops, not in select or switch
// "continue" statement begins the next iteration of the innermost enclosing "for" loop by advancing control to the end of the loop block.
func TestContinueStatement(t *testing.T) {
	//is continue statememt valid in normal flow - NO
	// t.Log("start testing")
	// continue //will fail with continue is not in a loop
	// t.Log("will this print")

	// is continue statement valid in switch - NO
	// switch i := 7; {
	// case i < 5:
	// 	continue // will fail with  continue is not in a loop
	// 	t.Logf("%d is < 5\n", i)
	// case i < 7:
	// 	t.Logf("%d is < 7\n", i)
	// default:
	// 	t.Logf("%d is >= 7\n", i)	
	// }
	common.Start("INFO: Unlike break, A continue statement is valid only in loops, not in select or switch. \n continue statement can work with labels.")
	common.End()

	// continue statement is valid only within loop. that too the loop should be in the same function
	common.Start("basic loop with continue")
	for i := 1; i <= 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				t.Log("i == j. skipping the rest of the body and going to next iteration")
				continue
				continue //this is never executed as the first continue skips it
			}
			t.Logf("i, j -> %d, %d\t", i, j)
		}
	}
	common.End()


	// continue statement is valid only within loop. that too the loop should be in the same function
	common.Start("skips odd number")
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			continue
		}
		t.Log("prints only even no:", i)
	}
	common.End()
}

// continue statement can work with labels
func TestContinueWithLabels(t *testing.T) {
	common.Start("continue with labels")
	start:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == j {
				t.Log("i == j. skipping the rest of the body and going to next itertion referenced by label start")
				continue start
			}
			t.Logf("i, j -> %d, %d\t", i, j)
		}
	}
	common.End()
}