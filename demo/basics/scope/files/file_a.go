// Package Scope demonstrates the scope of names is go.
//  - Package Scope: All top level identifiers are visible throughout the package. Exported top level identifiers are visible to other packages which imports it.
//  - Function Scope: Visible only within the function.
//  - Block Scope: Visible only within the block.
//  - Global Scope: No true global scope, but exported identifiers are visible in other packages.
package scope

// A identifier should start with Capital case to qualify as exported.
// Exported top level identifiers are visible to this entire package
// and to other packages which imports this package.
var ExportedPackageVariable string = "1.1 ExportedPackageVariable: I am visible throughout the entire package and to other packages which imports this package"

// Identifiers start with small case are not Exported.
// It is only visible to this package
var packageVariable string = "1.2 packageVariable: I am visible to this package. Not visible to other packages"

// identifiers start with small case are not Exported.
// It is only visible to this package
func packageFunction() {
    println("1.3. i am package Function, visible throughout this package") // println - a built in function
}


