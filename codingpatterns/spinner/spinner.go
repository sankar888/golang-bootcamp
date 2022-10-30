package main

import (
	"fmt"
	"time"
)

func main() {
	indefiniteSpinner()
}

func indefiniteSpinner() {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(1 * time.Second)
		}
	}
}
