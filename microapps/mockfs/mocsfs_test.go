package mockfs

import (
	"testing"
)

func TestPwd(t *testing.T) {
	var fs *Fs = NewMockFs()
	path := fs.Pwd()
	t.Log(path)
}

func TestRoot(t *testing.T) {
	var fs *Fs = NewMockFs()
	root := fs.Root()
	flag := root.IsDir()
	t.Log("root is dir", flag)
}
