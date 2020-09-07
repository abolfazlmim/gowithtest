// Golang program to illustrate the usage of 
// io.Copy() function 

// Including main package 
package main 

// Importing fmt, io, os, and strings 
import ( 
	"fmt"
	"io"
	"os"
	"strings"
) 

// Calling main 
func main() { 

	// Defining source 
	src := strings.NewReader("Nidhi: F\nRahul: M\nNisha: F\n") 

	// Defining destination using Stdout 
	dst := os.Stdout 

	// Calling Copy method with its parameters 
	bytes, err := io.Copy(dst, src) 

	// If error is not nil then panics 
	if err != nil { 
		panic(err) 
	} 

	// Prints output 
	fmt.Printf("The number of bytes are: %d\n", bytes) 
} 

