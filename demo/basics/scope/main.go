package main

import (
    "fmt"
    "github.com/sankar888/golang-bootcamp/demo/basics/scope/files" // In Go, packages are imported by fully qualified path of the package. if the import path / dir has two different package , it will result in error
     util "github.com/sankar888/golang-bootcamp/demo/common"
)


func main() {
    util.Start("Scope in Go:")
    fmt.Println("1. scope.ExportedPackageVariable from main package:", scope.ExportedPackageVariable)
    fmt.Println("")
    scope.PackageVarDemoFn()
    scope.PackageFunctionDemo()
    fmt.Println("")
    functionScopeDemo()
    util.End()
}


var i int = 0 // a package scoped identifier visible to all functions in this package

// functionScopeDemo function demonstrates the function dn block level scope
// In addition to the package scope, identifiers are scoped based on the function and blocks they are defined in.
// A package scoped identifier is visible to all functions in that package
// if there are any identifiers declared within the function with the same identifier name, it takes precedence
func functionScopeDemo() {
    i := 10 // shadows the package variable i = 0
    // the value of i would be 10. Identifiers are also scoped by the function and blocks they are defined it
    fmt.Println("2. functionScopeDemo: what would be the value of i inside function functionScopeDemo:", i)

    // Identifiers defined within block is visible only within that block. it shadows function and variable scope identifiers
    // A block scope can be created just by using a block {} or some statements like if, for, switch also creates a block scope
    if i = 20; i > 0 {
        fmt.Println("3. Block scope: what would the value of i within the if block:", i) // i would be 20 as the block scope shadows the function scope
    }

    {
        var j int = 100
        fmt.Println("4. custom block scope: value of j is visible only inside this block", j)
    }
    // uncommenting the below line will cause error as j is not visible outside of the block it is defined
    //fmt.Println(j)
}
