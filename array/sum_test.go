package main

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 65}

	sum := Sum(numbers)
	fmt.Println(sum)
}

func TestSum(t *testing.T) {
	t.Run("collection of 5 number s", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given , %v", got, want, numbers)
		}
	})
	t.Run("collection of any size ", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21
		if got != want {
			t.Errorf("got %d want %d given , %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v  want %v", got, want)
	}
}
