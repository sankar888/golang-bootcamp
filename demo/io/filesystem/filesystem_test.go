// This package has test functions which demonstrates the golang's apis available for filesystem
package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

// package os has the basic file and directory abstraction
func TestBasicConcepts(t *testing.T) {
	//create a file and returns *os.File and error
	err := os.MkdirAll("C:/Users/sankaraa/work/tmp/", os.ModeDir)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("C:/Users/sankaraa/work/tmp/file_created_by_golang.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file.Name())
	defer file.Close()
}

func TestStatFunction(t *testing.T) {
	// check if file exists
	info, err := os.Stat("/a/b/c")
	if err != nil {
		t.Log("is err is of type os.ErrNotExists?", errors.Is(err, os.ErrNotExist))
		t.Logf("when it is err, what is info: %v\n", info)
		//t.Fatal(err)
	}

	//What if the path points to a directory. yes stat can work with directory
	info, err = os.Stat("C:/Users/sankaraa") //info is of type fs.FileInfo
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("BaseName: %s, Size: %dKB, lastmodifiedtime: %v, isdirectory: %v\n", info.Name(), info.Size()/1024, info.ModTime(), info.IsDir())

	//is fs.FileInfo got as result of Stat() call pointer? Yes
	t.Logf("fs of type %T, pointer %p\n", info, info) //info is a pointer
	callMe(info)

	//what does the stat of a file gives back
	info, err = os.Stat("C:/Users/sankaraa/.viminfo")
	if err != nil {
		t.Error(err)
	}
	t.Logf("BaseName: %s, Size: %dKB, lastmodifiedtime: %v, isdirectory: %v\n", info.Name(), info.Size()/1024, info.ModTime(), info.IsDir())
	t.Logf("Mode of a FileInfo %v\n", info.Mode())               //-rw-rw-rw-
	t.Logf("What would FileInfo Sys() returns %v\n", info.Sys()) //Don't know what is this. skipping for now.
}

func TestCreatingFileAndDirectory(t *testing.T) {
	//how to create a normal file.
	file, err := os.Create("C:/Users/sankaraa/work/tmp/file_created_by_golang.txt")
	if err != nil {
		t.Error(err)
	}
	//write some contents to the file
	writeFile(file, "Hai. Hello. How are u ?")
	err = file.Sync()
	if err != nil {
		t.Error(err)
	}
	printFileContents(file, t)
	defer file.Close()

	//what if we create the same file another time. the old file will be truncated and new file will be created
	file, err = os.Create("C:/Users/sankaraa/work/tmp/file_created_by_golang.txt")
	if err != nil {
		t.Error(err)
	}
	printFileContents(file, t)
	err = file.Close()
	if err != nil {
		t.Error(err)
	}
}

func callMe(info fs.FileInfo) {
	fmt.Printf("fs of type %T, pointer %p\n", info, info)
}

func writeFile(file *os.File, contents string) {
	//how to write to a file, using writer
	_, err := file.WriteString(contents)
	if err != nil {
		log.Fatal(err)
	}
}

func printFileContents(file *os.File, t *testing.T) {
	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		t.Error(err)
	}
	t.Logf("contents of file %s is %s\n", file.Name(), string(buffer[0:n]))
}
