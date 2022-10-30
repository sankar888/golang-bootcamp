package snippets

import (
	"testing"
)

func TestIsvalildPath(t *testing.T) {
	t.Log(t.Name())
	var path string = ".././abc/.. "
	if !IsvalidPath(path) {
		t.Log("test failed for input", path)
		t.Fail()
	}
}
