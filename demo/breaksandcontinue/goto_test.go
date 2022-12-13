package breaksandcontinue

import (
	"testing"
	"github.com/sankar888/golang-bootcamp/demo/common"
)

// Refer: https://go.dev/ref/spec#Goto_statements
//1. A "goto" statement transfers control to the statement with the corresponding label within the same function.
//2. Executing the "goto" statement must not cause any variables to come into scope that were not already in scope at the point of the goto. 
// For instance, this example:
// 	goto L  // BAD
// 	v := 3
// L:
// is erroneous because the jump to label L skips the creation of v.
//3. A "goto" statement outside a block cannot jump to a label inside that block.

func TestBasicgoto(t *testing.T) {
	common.Start("basic goto usage")
	for i := 1; i < 10; i++ {
		t.Logf("looping.. %d\n", i)
		if i % 2 == 0 {
			t.Logf("found even number %d. goto end\n", i)
			goto end
		}
	}
	end: t.Log("End.")
	common.End()

	// all defined labels should be used
	common.Start("All defined labels should be used. else error is thrown.")
	//start: 
	//t.Log("starting..") //will throuw error label start defined and not used
	common.End()
}


// this test will throw error, because labels are valid only within same function
func TestGotoScope(t *testing.T) {
	// loopUntil := func(until int) {
	// 	if until < 0 {
	// 		t.Log("couldn't count to negative number.")
	// 		t.Fail()
	// 	} 
	// 	i := 0
	// 	for {
	// 		if i == until {
	// 			goto end
	// 		}
	// 		t.Logf("looping... %d\n", i)
	// 		i++
	// 	}
	// }

	// end: t.Log("End.")
	// loopUntil(3)
}

func TestGotoVariablScope(t *testing.T) {
	common.Start("goto and new variables")
	t.Log("Executing the goto statement must not cause any variables to come into scope that were not already in scope at the point of the goto.")
	animal := "dog"	
	if animal == "dog" {
		goto dog //this goto will work, why? don't know.
		leg := 4
		t.Logf("dog has %d legs.\n", leg)
	}

	// goto dog //this will fails as goto skips the creation of leg in the same scope
	// leg := 4
	// t.Logf("dog has %d legs.\n", leg)
	dog: t.Log("Dog barks..")
	common.End()
}


