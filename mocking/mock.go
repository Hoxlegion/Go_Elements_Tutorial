package main

import (
	"fmt"
	"io"
	"os"
	"time"
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
    sleeper := &ConfigureableSleeper{1 * time.Second, time.Sleep}
    CountDown(os.Stdout, sleeper)
}



