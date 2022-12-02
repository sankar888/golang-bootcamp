package learnjson

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestBasicJsonMarshalling(t *testing.T) {
	//encode an int
	var id int = 1984
	b, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("encoded json int ", string(b))

	//encode an array
	var arr [3]string = [3]string{"maa", "palaa", "vaalai"}
	b, err = json.Marshal(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json representation of array ", string(b))

	//encode an map
	var countMap map[string]int = map[string]int{}
	countMap["hai"] = 3
	countMap["hello"] = 6
	b, err = json.Marshal(countMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json representation of map", string(b))

	//encode a struct
	alarm := struct {
		Timestamp time.Time
		Message   string
	}{
		Timestamp: time.Now(),
		Message:   "Wake up, its time to run \U0001F60A",
	}
	b, err = json.Marshal(alarm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json representation of struct", string(b))

	//encode a function
	add := func(a int, b int) int {
		return a + b
	}
	b, err = json.Marshal(add)
	if err == nil {
		t.Logf("Json couldn't marshal function types. expected err.")
		t.Fail()
	}
	fmt.Println(err)
	fmt.Println("json representation of empty struct", string(b))
}
