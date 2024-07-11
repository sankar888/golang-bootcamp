// Package Comments explains comments in Go and how they are used for Documentation.
// There are two types of comments: single-line and multi-line comments.
//  - Single-line comments start with // and extend to the end of the line.
//  - Multi Line Comments starts and ends with /* */ respectively and can span multiple lines
//
// Go documentation tool interprets comments placed immediately before top-level package, function, type, or variable declarations as documentation
// These comments are used by the go doc tool to generate documentation.
// Documentation comments typically start with the name of the entity being documented.
//
// Usually lengthy top level package documentation with examples uses /**/ multi-line documentation. Refer documentation of gob package https://cs.opensource.google/go/go/+/go1.22.5:src/encoding/gob/doc.go
//
// https://tip.golang.org/doc/comment describes about how to write go documentation comments in detail
package comments

import (
    "fmt"
)

// One is a constant representing numeric one
const One = 1

var Language string = "English" // Language is a variable representing the English language

// country is a not exported name and documentation will not be generated for it
var country string = "India"

// Constants can be documented as a whole or individually
const (
    Tens        = 10 // Constant for 10
    Hundreds    = 100 // Constant denoting 100
    Thousands   = 1000
)

// HelloWorld prints "Hello World!"
func HelloWorld() {
    fmt.Println("hello World!")
}

/*
Hello greets a person with the given name.
*/
func Hello(name string) {
    fmt.Println("Hello", name)
}

// ComplexFunction demonstrate which comments are not treated as documentation comments.
// The source code of this function is as follows
//  func ComplexFunction(value interface{}) {
//      // This is not a documentation comment, because it is not at top level
//       var aPoem string = "Do not go gentle into that good night, Old age should burn and rave at close of day; Rage, rage against the dying of the light."
//   }
func ComplexFunction(value interface{}) {
    // This is not a documentation comment, because it is not at top level
    var aPoem string = "Do not go gentle into that good night, Old age should burn and rave at close of day; Rage, rage against the dying of the light."
}

// Two widely used documentation tools of go are
//  - go doc - SubCommand of go utility. Serves manual or help for any go Name. prints the documentation in terminal
//  - godoc - Tool which generates and serves documentation in html format
var GoDocTools []string = []string{"go doc", "godoc"}
