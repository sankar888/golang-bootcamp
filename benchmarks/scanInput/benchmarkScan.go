package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(scanBigString())
	elapsed := time.Since(start)
	fmt.Printf("scan took %s \n", elapsed)

	start = time.Now()
	fmt.Println(readBigString())
	elapsed = time.Since(start)
	fmt.Printf("read took %s \n", elapsed)
}

func scanBigString() string {
	var s string
	fmt.Scanln(&s)
	return s
}

func readBigString() string {
	rdr := bufio.NewReader(os.Stdin)
	tmp, _ := rdr.ReadString('\n')
	tmp = strings.TrimRight(tmp, "\r\n")
	return tmp

}
