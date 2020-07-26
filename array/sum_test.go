package main

import (
	"fmt"
	"testing"
)

func ExampleSum() {
	numbers := [5]int{1, 2, 3, 4, 65}

	sum := Sum(numbers)
	fmt.Println(sum)
}



func TestSum(t *testing.T) {
	t.Run("collection of 5 number s", func(t *testing.T){
		numbers :=[]int
	})
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("got %d want %d given , %v", got, want, numbers)
	}
}
