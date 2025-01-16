package main

import (
	"fmt"
	"io"
	"os"
)

const (
    finalWord = "Go!"
    countdownStart = 3
)

func CountDown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        sleeper.Sleep()
    }

    for i := countdownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
    }

    fmt.Fprint(out, finalWord)
}

func main() {
    sleeper := &DefaultSleeper{}
    CountDown(os.Stdout, sleeper)
}



