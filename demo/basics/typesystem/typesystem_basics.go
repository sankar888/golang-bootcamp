// Package typesystem demonstrates the go type system.
// GO is a strictly typed language, meaning every object / value has a type assigned to it
// and the type determines what operations can be applied to that value
package typesystem

import (
    "fmt"
    "reflect"
)
// GO is a strictly typed language, meaning every object / value has a type assigned to it
// and the type determines what operations can be applied to that value.
//
// Type of a value or object can be obtained in go by using
//  - reflect package
//  - fmt package formatter
// Example:
//  var a string = "string"
//  // reflect.type has many useful functions which enables us to investigate the methods and other aspects of a type
//  var t reflect.Type = reflect.TypeOf(a)
//
//  // simple way which returns the string representation of type
//  fmt.Printf("%T", a)
const Overview = "Overview"

// KnowYourType function takes in any value
// and returns the string representation of type and kind of any value
// it uses reflect api to get the type of value
func KnowYourType(value interface{}) (t, k string) {
    t0 := reflect.TypeOf(value)
    return t0.String(), t0.Kind().String()
}

// TypeOf function takes in any value
// and returns the string representation of type as returned by "fmt" formatter
func Typeof(value interface{}) string {
    return fmt.Sprintf("%T", value)
}