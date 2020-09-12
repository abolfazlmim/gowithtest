package main

import (
	"fmt"
)

func main() {
	q := [...]int{5, 6, 7, 8, 9, 10}
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	for i, _ := range a {
		fmt.Printf("%d \n", i)
	}
	for i, v := range q {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Printf("%T\n", q)
}
