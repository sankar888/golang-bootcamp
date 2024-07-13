package typesystem

import (
    "testing"
)

// TestKnowYourType demonstrates the type of various values
// look at the various test cases and understand the type of various values
func TestKnowYourType(test *testing.T) {
    type TCase struct {
        Value interface{}
        OType string
        OKind string
    }

    var i8 uint8
    b := true
    var s string = "string"
    function := func(a int) {}
    var j int = 123
    k := &j

    tcases := []TCase {
        {
            Value: 0,
            OType: "int",
            OKind: "int",
        },
        {
            Value: i8,
            OType: "uint8",
            OKind: "uint8",
        },
        {
            Value: int32(-10),
            OType: "int32",
            OKind: "int32",
        },
        {
            Value: 34.45,
            OType: "float64",
            OKind: "float64",
        },
        {
            Value: float32(23.45),
            OType: "float32",
            OKind: "float32",
        },
        {
            Value: b,
            OType: "bool",
            OKind: "bool",
        },
        {
            Value: s,
            OType: "string",
            OKind: "string",
        },
        {
            Value: function,
            OType: "func(int)",
            OKind: "func",
        },
        {
            Value: [2]int{1,2},
            OType: "[2]int",
            OKind: "array",
        },
        {
            Value: []int{1,2},
            OType: "[]int",
            OKind: "slice",
        },
        {
            Value: map[string]int{"one":1, "two":2},
            OType: "map[string]int",
            OKind: "map",
        },
        {
            Value: Animal{},
            OType: "typesystem.Animal",
            OKind: "struct",
        },
        {
            Value: &j,
            OType: "*int",
            OKind: "ptr",
        },
        {
            Value: k,
            OType: "*int",
            OKind: "ptr",
        },
    }
    for _, tcase := range tcases {
        if t, k := KnowYourType(tcase.Value); t != tcase.OType || k != tcase.OKind {
            test.Errorf("case: %v failed. got Type %s, Kind: %s\n", tcase, t, k)
        }
    }
}

type Animal struct {
}



