package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

    checkSum := func(t testing.TB, result, want int, numbers []int) {
        t.Helper()

        if result != want {
            t.Errorf("got %d want %d given, %v", want, result, numbers)
        }
    }
    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1, 2, 3, 4, 5}


        result := Sum(numbers)
        want := 15

        checkSum(t, result, want, numbers)
    })

    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        result := Sum(numbers)
        want := 6

        checkSum(t, result, want, numbers)
    })
}

func TestSumAll(t *testing.T) {
    result := SumAll([]int{1,2}, []int{0,9})
    want := []int{3, 9}

    if !slices.Equal[[]int](result, want) {
        t.Errorf("got %d want %d", result, want)

    }
}

func TestSumAllTails(t *testing.T) {

    checkSums := func(t testing.TB, result, want []int) {
        t.Helper()

        if !slices.Equal[[]int](result, want) {
            t.Errorf("got %v want %v", result, want)
        }
    }
    t.Run("make the sums of some slices", func(t *testing.T) {
        result := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}

        checkSums(t, result, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        result := SumAllTails([]int{}, []int{0, 9})
        want := []int{0, 9}

        checkSums(t, result, want)
    })
}

