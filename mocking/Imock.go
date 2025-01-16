package main

import "time"

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}

type SpyTime struct {
	durationSlept time.Duration
}

type SpySleeper struct {
	Calls int
}

type SpyCountDownOperations struct {
	Calls []string
}

type ConfigureableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (e int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *ConfigureableSleeper) Sleep() {
    c.sleep(c.duration)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
