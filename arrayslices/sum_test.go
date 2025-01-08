package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}


        got := Sum(numbers)
        want := 15

        log(t, got, want, numbers)
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        log(t, got, want, numbers)
    })
}

func TestSumAll(t *testing.T) {
    result := SumAll([]int{1,2}, []int{0,9})
    want := []int{3, 9}

    if !slices.Equal[[]int](result, want) {
        t.Errorf("got %d want %d", result, want)

    }

}

func log(t *testing.T, got, want int, numbers []int) {
        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
}
