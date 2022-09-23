package common

import "fmt"

func Start(name string) {
	fmt.Printf("**%v**\n", name)
}

func End() {
	fmt.Println("---end---")
	fmt.Println("")
}
