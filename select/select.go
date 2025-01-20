package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
    aDuration := measureReponseTime(a)
    bDuration := measureReponseTime(b)

    if aDuration < bDuration {
        return a
    }

    return b
}

func measureReponseTime(url string) time.Duration {
    start := time.Now()
    http.Get(url)
    return time.Since(start)
}



