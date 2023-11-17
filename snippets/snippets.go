package snippets

import (
	"bytes"
	"errors"
	"log"
	"regexp"
	"strings"
	"text/template"
)

// simplify path takes in a filesystem path which can be absolute or relative and
// can have . and ..
// can have multiple /
// and cleans and simplify the path
// Ex: .././../////abc.///./../abv/ -> ../../abv
func simplifyPath(path string) (cleanedPath string, err error) {
	path = strings.TrimSpace(path)
	if length := len(path); length == 0 {
		return "", errors.New("Got Empty Path. Expecting a path of non zero length")
	}
	//what if the path itself is not valid
	//..../../..../dbc/../sc/// - not valid conatins four dot.

	return "", nil
}

// consider the pathis not empty
func IsvalidPath(path string) bool {
	//..../../..../dbc/../sc///
	//A path is valid if it only contains . or .. or / or [a-z][A-Z][0-9] regex: [\.\/a-zA-Z0-9]+
	//Also path . or .. should be followed by /

	//or we could chek for invalid path
	//anything which has character other than [a-z][A-Z][0-9], . and .. is invalid
	//. or .. should be followed only by / and not by [a-z][A-Z][0-9]
	//
	exp := regexp.MustCompile(`[^\.\/a-zA-Z0-9]|\.{3,}|\.[a-zA-Z0-9]`)
	return !exp.MatchString(path)
}

// Simplify the
func Simplify(path string) string {
	//..///abc/./a//xyz/../..
	reg := regexp.MustCompile(`\/+`)
	path = reg.ReplaceAllString(path, "/")
	var parts = strings.Split(path, "/")

	l := len(parts)
	cparts := make([]string, l)
	cpindex := 0
	for _, c := range parts {
		if c == "." {
			continue
		}
		if c == ".." && cpindex > 0 && cparts[cpindex] != ".." {
			cpindex--
			continue
		}
		cparts[cpindex] = c
		cpindex++
	}
	cparts = cparts[:cpindex]
	return strings.Join(cparts, "/")
}

func RenderTemplate(templateStr string, data any) string {
	t := template.Must(template.New("tmp").Parse(templateStr))
	var writer bytes.Buffer
	if err := t.Execute(&writer, data); err != nil {
		log.Panic(err)
	}
	return writer.String()
}
