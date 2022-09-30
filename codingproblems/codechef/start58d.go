package main

import (
	"fmt"
	"log"
)

func main() {
	//reachTarget()
	//rankListPage()
	removeBadElements()
}

// https://www.codechef.com/START58D/problems/REACHTARGET
func reachTarget() {
	var t int
	_, err := fmt.Scan(&t)
	if err != nil {
		log.Fatal(err)
	}

	for ; t > 0; t-- {
		var x, y int
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(x - y)
	}
}

// https://www.codechef.com/START58D/problems/RANKLISTPAGE
func rankListPage() {
	var t int
	_, err := fmt.Scan(&t)
	if err != nil {
		log.Fatal(err)
	}

	for ; t > 0; t-- {
		var x int
		_, err := fmt.Scan(&x)
		if err != nil {
			log.Fatal(err)
		}

		var page int
		switch {
		case x <= 25:
			page = 1
		case x%25 == 0:
			{
				page = (x / 25)
			}
		default:
			{
				page = (x / 25) + 1
			}
		}
		fmt.Println(page)
	}
}

func removeBadElements() {
	var t int
	_, err := fmt.Scan(&t)
	if err != nil {
		log.Fatal(err)
	}

	for ; t > 0; t-- {
		var n int
		fmt.Scan(&n)
		var arr []int = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}
		fmt.Printf("%v. %v \n", n, arr)
	}
}
