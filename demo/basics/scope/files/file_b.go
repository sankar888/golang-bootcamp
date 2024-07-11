package scope

import (
    "fmt"
)

// PackageVarDemoFn demonstrates the scope of package variable and Exported package variable
func PackageVarDemoFn() {
    fmt.Println(packageVariable) //there is no need for import statement among different files of a same package
    fmt.Println(ExportedPackageVariable)
}

// PackageFunctionDemo function demonstrates the scope of package functions
// function packageFunction should be visible to this file without any import as both files declare the same package
func PackageFunctionDemo() {
    packageFunction()
}