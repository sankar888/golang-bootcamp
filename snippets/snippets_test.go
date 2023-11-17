package snippets

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

/*
 * A Function to generate sql based on the string template
 * given a array of struct, to the string template, will give templated string
 * can be achieved using simeple printf or could use the golang text/template
 * text/template - https://pkg.go.dev/text/template@go1.20.7
 */

func TestQuotesTemplate(t *testing.T) {
	templateStr := `{{- range . -}}
	(Quote: {{ .Quote }}),
	{{- end -}}`
	data := []struct {
		Quote string
	}{
		{Quote: "quote1"},
		{Quote: "quote2"},
		{Quote: "quote3"},
	}
	t.Log(RenderTemplate(templateStr, data))
}

func TestParseCsv(t *testing.T) {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("Error opening CSV file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file or error
		}
		fmt.Println("Row:", row)
	}
}

type test struct {
	module       string
	testClass    string
	testName     string
	timeTakenSec string
}

// A function to walk through embedded test xml report files, parse it and get the info
func TestParseXMl(t *testing.T) {

	type TestCase struct {
		Name      string `xml:"name,attr"`
		Time      string `xml:"time,attr"`
		ClassName string `xml:"classname,attr"`
	}

	type testSuite struct {
		XMLName   xml.Name   `xml:"testsuite"`
		TestCases []TestCase `xml:"testcase"`
	}

	rootDirectory := "C:/Users/sankaraa/work/tmp/archive"
	output := "C:/Users/sankaraa/work/tmp/embeddedTests.csv"
	var tests []test = make([]test, 0)

	processFile := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".xml") {
			path = filepath.ToSlash(path)
			parts := strings.Split(path, "/")
			if l := len(parts); l > 4 {
				module := parts[l-4]

				content, err := os.ReadFile(path)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return nil
				}

				var ts testSuite
				xmlContent := string(content)
				err = xml.Unmarshal([]byte(xmlContent), &ts)
				if err != nil {
					fmt.Println("Error parsing XML:", err)
					return nil
				}

				// Iterate through the test cases and print the name and time
				for _, testCase := range ts.TestCases {
					var tst test = test{}
					tst.module = module
					tst.testClass = testCase.ClassName
					tst.testName = testCase.Name
					tst.timeTakenSec = testCase.Time
					tests = append(tests, tst)
					fmt.Println(tst)
				}
			}
		}
		return nil
	}

	// Start walking through the directory and its subdirectories
	err := filepath.Walk(rootDirectory, processFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	writeTestsCsv(tests, output)
}

func writeTestsCsv(tests []test, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"Module", "TestClass", "TestName", "TimeTakenSec"}
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	// Write data rows
	for _, t := range tests {
		row := []string{t.module, t.testClass, t.testName, t.timeTakenSec}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}
