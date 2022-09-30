package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	common.Start("**Bold | Strikethrough | Underline**")
	applyStyle(Bold | Strikethrough | Underline)
	common.End()

	common.Start("**Bold**")
	applyStyle(Bold)
	common.End()

	common.Start("**Unknown 123**")
	applyStyle(123) // This would produce a invalid style, if the client systems use it improper
	common.End()
}

// Style is used to style the text
// Multiple styles can be applied at once
// We could use 1,2,3,4.. as constant values,
// but if we choose to use bit for flags and multiple flags can be set at once
// we need to crepresent each flag by power of 2, which gives a unique bit for each flag
const (
	Bold          int = 1 //binary 00001
	Italic        int = 2 //binary 00010
	Underline     int = 4 //binary 00100
	Strikethrough int = 8 //binary 01000
)

// applyStyle applies the styles set by flag
// The flag can represent more than one style
func applyStyle(styleFlags int) {
	if styleFlags&Bold == Bold {
		fmt.Println("Bold is set")
	}

	if styleFlags&Italic == Italic {
		fmt.Println("Italic is set")
	}

	if styleFlags&Underline == Underline {
		fmt.Println("Underline is set")
	}

	if styleFlags&Strikethrough == Strikethrough {
		fmt.Println("Strikethrough is set")
	}
}
