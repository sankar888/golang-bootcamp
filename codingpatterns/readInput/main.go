/*
This package demonstrates how to read input from terminal and string using fmt package
*/
package main

import (
	"fmt"

	"github.com/sankar888/golang-bootcamp/demo/common"
)

func main() {
	//printEmoticon()
	//fmtDemoScan()
	//fmtDemoScanf()
	//fmtDemoScanln()
	stringversionOfScanFunctions()
}

// fmt package provides functions to scan / read data from std.in or from any reader
// Scan, Fscan, Sscan treat newlines in the input as spaces.
// Scanln, Fscanln and Sscanln stop scanning at a newline and require that the items be followed by a newline or EOF.
// Scanf, Fscanf, and Sscanf parse the arguments according to a format string, analogous to that of Printf. In the text that follows, 'space' means any Unicode whitespace character except newline.
func fmtDemoScan() {
	common.Start("fmt scan usage")
	//basic scan, without formatting or newline
	var a, b, c int
	fmt.Println("Enter values for integers a b and c separated by space or newline:")
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println("scanned values of a, b, c:", a, b, c)
	common.End()
}

// scan formatted input
func fmtDemoScanf() {
	common.Start("formatted scan using fmt.Scanf()")
	fmt.Println("Enter a expression of format a + b = c")
	var a, b, c int
	var tUp, tDown rune = '\U0001F44D', '\U0001F44E'
	_, err := fmt.Scanf("%d + %d = %d", &a, &b, &c)
	if err != nil {
		panic(err)
	}
	if a+b == c {
		fmt.Printf("You are correct %c\n %d + %d = %d\n", tUp, a, b, c)
	} else {
		fmt.Printf("You are Wrong %c\n %d + %d != %d\n", tDown, a, b, c)
	}
	common.End()
}

// scans a input params from std.in. space separated arguments are scanned
// unlink fmt.Scan() fn, new line are not considered as spaces
// A single scan cannot span across multiple lines
func fmtDemoScanln() {
	common.Start("scan line function demo")
	fmt.Println("Enter atleast three space separated words")
	var a, b, c string
	fmt.Scanln(&a, &b, &c)
	fmt.Println("Values of a, b, c scanned", a, b, c)
	common.End()
}

func stringversionOfScanFunctions() {
	common.Start("Scan from String Demo")
	var p1, p2 string
	fmt.Sscanf("Jack and Jerry went up the hill", "%s and %s", &p1, &p2)
	fmt.Println("The values of p1 p2 :", p1, p2)

	var s1, s2, s3, s4 string
	fmt.Sscan(
		`hai
		how r
		you`, &s1, &s2, &s3, &s4)
	fmt.Println("values of scanned s1 s2 s3 s4:", s1, s2, s3, s4)
	common.End()
}

// In Golang utf characters are represented by rune a int32 alias.
func printEmoticon() {
	var emoji rune = '\U00000000' //unicode characters or literals ar represented by \U<8 digit code> or \u<4digit code>
	var shortEmoji rune = '\u0B85'
	fmt.Printf("%c \n", shortEmoji)
	var emojis []rune = make([]rune, 0)
	for i := 0; i < 1000; i++ {
		emoji += 1
		emojis = append(emojis, emoji)
	}
	fmt.Printf("%c \t", emojis)
}
