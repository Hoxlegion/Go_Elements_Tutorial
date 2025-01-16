package main

import "time"

type Sleeper interface {
    Sleep()
}

type SpySleeper struct {
    Calls int
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

