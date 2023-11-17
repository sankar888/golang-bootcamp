// Package mockfs defines a mockfs. It is an inmemory filesystem
package mockfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"json"
	"regexp"
	"strings"
	"time"
)

const invalidPath string = `[^\.\/a-zA-Z0-9]|\.{3,}|\.[a-zA-Z0-9]`

var invalidPathRegex *regexp.Regexp = regexp.MustCompile(invalidPath)

// File represents a file in mock file system. The attributes are not exposed outside
type File struct {
	name     string
	content  []byte
	meta     metadata
	children []*File //only valid for directory
	dir      bool
	parent   *File
}

// metadata represents the meta information of the file
type metadata struct {
	createdOn      time.Time
	lastModifiedOn time.Time
}

type Fs struct {
	root *File
	cwd  *File
}

// NewMockFs function creates a mck file system
func NewMockFs() *Fs {
	root := File{
		name:     "/",
		children: make([]*File, 0),
		dir:      true,
		meta:     getCreationMeta(),
		parent:   nil,
	}
	return &Fs{
		root: &root,
		cwd:  &root,
	}
}

func getCreationMeta() metadata {
	now := time.Now()
	meta := metadata{
		createdOn:      now,
		lastModifiedOn: now,
	}
	return meta
}

/*
	func (f *Fs) CreateFile(path string, content []byte) (created bool, err error) {
		//validate path, get the dir node which represents the path
		//create file node and add it to the directory
		path = strings.TrimSpace(path)
		if l := len(path); l == 0 {
			return false, errors.New("Could not create file with Empty name")
		}
		//TODO: define and check illega characters for a file name

}

func (f *Fs) Rm(path string) error {

}
*/

func (f *Fs) CreateDir(path string) error {
	if l := len(path); l == 0 {
		return errors.New("the path to create a directory should not be blank")
	}
	if (path == "." || )
}

func (f *Fs) CurrentDir() (dir *File) {
	return f.cwd
}

func (f *Fs) Pwd() (path string) {
	//get current directory, get its name, if it has parent, get its parent concat with earlier name
	//untill root
	var parts []string = make([]string, 0)
	cdir := f.CurrentDir()
	if cdir.name == "/" {
		return "/"
	}
	for {
		parts = append(parts, cdir.name)
		if cdir = cdir.parent; cdir.name == "/" {
			break
		}
	}
	//c b a -> a b c
	fmt.Println("parts ", parts)
	var partStr []byte = make([]byte, 0)
	partStr = append(partStr, '/')
	for i := len(parts) - 1; i >= 0; i++ {
		partStr = append(partStr, parts[i]...)
		partStr = append(partStr, '/')
	}
	partStr = partStr[:len(partStr)-1]
	fmt.Println("pwd path", partStr)
	return string(partStr)
}

/*
func (f *Fs) ListDirContents(path string) (entries []*File, err error) {

}

func (f *Fs) ChangeWorkingDir(path string) error {

}
*/

func (fs *Fs) Root() *File {
	return fs.root
}

func (f *File) IsDir() bool {
	return f.dir && f.children != nil
}

/*
func (f *Fs) getDir(path string) (entry File, err error) {

	parts, err := getPathParts(path)
	if err != nil {
		return File{}, err
	}

	//if the path is absolute traverse from root
	if isAbsolutePath(parts) {

	} else {
		//if the path is relative, traverse from current wdir

	}

	//how to traverse an entry A.
	//check if the entry A is dir. get child contents of dir A. check the directory matching the path part. and go on till u reach the target

}
*/

func (f *Fs) addFiletoDir(dir File, file ...File) {

}

func isAbsolutePath(parts []string) bool {
	//if the path starts with / it is absolute
	//if the path starts with anything other than / or . or .., it is relative
	return parts[0] == "/"
}

func isvalidPath(path string) bool {
	return !invalidPathRegex.MatchString(path)
}

func getPathParts(path string) (parts []string, err error) {
	if !isvalidPath(path) {
		return nil, errors.New(fmt.Sprintf("%s is not a valid mockfs path", path))
	}
	reg := regexp.MustCompile(`\/+`)
	path = reg.ReplaceAllString(path, "/")
	rparts := strings.Split(path, "/")

	l := len(rparts)
	parts = make([]string, l)
	pindex := 0
	for _, c := range parts {
		if c == "." {
			continue
		}
		if c == ".." && pindex > 0 && parts[pindex] != ".." {
			pindex--
			continue
		}
		parts[pindex] = c
		pindex++
	}
	if parts[0] == "" {
		parts[0] = "/"
	}
	parts = parts[:pindex]
	return parts, nil
}
