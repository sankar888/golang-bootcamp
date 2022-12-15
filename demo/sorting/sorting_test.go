// Package learnsorting demonstrates, how to use sort pckage in golang
// sort package has a list of sort methods for
// 1. Ints
// 2. strings
// 3. slice
// 4. any type which implements sort.Interface
// sort packages exposes two types of sorting function
// 1. Normal - non determininstic, doesn't mintain the original order of equal elements.
// 2. Stable sort - is deterministic, it maintains the original order of equal elements.
// sort package also has search functions to search a slice for an element
package learnsorting

import (
	"sort"
	"testing"
	"github.com/sankar888/golang-bootcamp/demo/common"
)

// This test demonstrates the usage of 
// func Ints(x []int)
//   Ints sorts a slice of ints in increasing order.
// func IntsAreSorted(x []int) bool
//   IntsAreSorted reports whether the slice x is sorted in increasing order.
func TestSortInts(t *testing.T) {
	common.Start("sorting slice of ints")
	var ints []int = []int{23,-98, 20, 0, 3, 1}
	sort.Ints(ints)
	t.Log("sorted slice", ints)
	common.End()

	common.Start("sorting empty slice")
	ints = []int{}
	sort.Ints(ints)
	t.Log("sorted empty slice", ints)
	common.End()

	common.Start("test if a slice of ints is already sorted")
	ints = []int{1, 2, 2, 3, 4, 91}
	sorted := sort.IntsAreSorted(ints)
	t.Logf("is %v in ascending order? %v\n", ints, sorted)
	common.End()

	common.Start("sort un-initiated slice")
	var slice []int //the default value is empty slice
	t.Log(slice)
	sort.Ints(slice) 
	t.Log("sorted uninitiated slice", slice)
	common.End()

	common.Start("can int8 be sorted using sort.Ints")
	//var int8s []int8 = []int8{1, 9, 0, -2} //No it will not work. []int8 can't be used in place of []int. 
	// wither we could loop and convert one by one or use unsafe package to do conversion
	//sort.Ints(ints)
	//t.Log("sorted slice of int8", int8s)
	common.End()
}

// This test demonstrates the use of sorting a slice of float64
// func Float64s(x []float64)
// func Float64sAreSorted(x []float64) bool
func TestSortFloat64s(t *testing.T) {
	common.Start("sort an slice of float")
	var floats []float64 = []float64{1.2, 0, -2, 4.6, 9.3, 9.5}	
	sort.Float64s(floats)
	t.Log("sorted float slice", floats)
	floats = []float64{9, 3, 5, 6, 1} // are treated as float literals
	sort.Float64s(floats)
	t.Log("sorted float slice", floats)
	t.Logf("is %v in ascending order? %v\n", floats, sort.Float64sAreSorted(floats))
	common.End()
}


func TestSortingAnyTypeOfSlice() {

}	

//TODO: sorting any slice, les function
//TODO: sort.Interface
//TODO: sorting of strings
//TODO: search methods in sort package




