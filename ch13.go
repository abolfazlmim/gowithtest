package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "----"))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0])
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("index:", i, "value: ", os.Args[i])
	}
}
