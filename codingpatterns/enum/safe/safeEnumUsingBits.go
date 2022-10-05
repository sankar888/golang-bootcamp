package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	common.Start("**Bold, Italic and Underline**")
	applySafeStyle(B | I | U)
	common.End()

	common.Start("**Bold and Underline**")
	applySafeStyle(B | U)
	common.End()

	common.Start("**Unknown 123**")
	applySafeStyle(99) // panic: Invalid Style Combination
	common.End()
}

type Style int8

// Style is used to style the text
// Multiple styles can be applied at once
// We could use 1,2,3,4.. as constant values,
// we need to crepresent each flag by power of 2, which gives a unique bit for each flag
// Custom Type approach not suitable when flags are more than 3
const (
	B   Style = 1
	I   Style = 2
	U   Style = 4
	BI  Style = B | I
	BU  Style = B | U
	IU  Style = I | U
	BIU Style = B | I | U
)

// applyStyle applies the styles set by flag
// The flag can represent more than one style
func applySafeStyle(flag Style) {
	if flag < 0 || flag > BIU {
		panic("Invalid Style Combination")
	}
	if flag&B == B {
		fmt.Println("Bold is set")
	}

	if flag&I == I {
		fmt.Println("Italic is set")
	}

	if flag&U == U {
		fmt.Println("Underline is set")
	}

}
