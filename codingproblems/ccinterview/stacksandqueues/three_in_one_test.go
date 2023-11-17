// cracking the coding interview problems
// chapter 3: stacks and queues
package stacksandqueues

// 3.1 Describe how you could use a single array to implement 3 stacks ?
// Hints: #2, #12, #38, #58
const Question = 3.1

// soln: strategy 1: allocate continious area for stcks seperated by buffer no of values
// strategy 2: each stack next position is fixed, s1 index -> 0, 3, 6 , s2 index -> 1, 4, 7 s3 index -> 2, 5, 8, ...
// selecting strategy 2
const Soln = 3.1

// A unique identifier which identifies stack
// All functions accepts this identifier to determine which stack to operate on
type StackIdentifier uint

const (
	StackOne   StackIdentifier = 0
	StackTwo   StackIdentifier = 1
	StackThree StackIdentifier = 2
)

// ThreeStack holds the array and the offset positions of three stacks
type ThreeStack struct {
}

func newThreeStack() ThreeStack {

}

func push()

// what is the difference between generics and interface{} - any in go ?
// why gnenerics hsould be qualitied in golang ?
// when to use generics , when to use interface{}
