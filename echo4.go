package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Ommit trailing new line")
var sep = flag.String("s", " ", "seprator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
