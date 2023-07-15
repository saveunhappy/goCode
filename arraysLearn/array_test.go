package main

import (
	"reflect"
	"testing"
)

//func TestSum(t *testing.T) {
//
//	numbers := [5]int{1, 2, 3, 4, 5}
//
//	got := Sum(numbers)
//	want := 15
//
//	if got != want {
//		t.Errorf("got %d want %d given, %v", got, want, numbers)
//	}
//}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

//func Sum(numbers [5]int) int {
//	sum := 0
//	for i := 0; i < 5; i++ {
//		sum += numbers[i]
//	}
//	return sum
//}
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func SumAll(numbersToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersToSum)
	sums := make([]int, len(numbersToSum))

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
