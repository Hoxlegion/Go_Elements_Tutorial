package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}

		}
	}
}

func main() {
	sleeper := &ConfigureableSleeper{1 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}


