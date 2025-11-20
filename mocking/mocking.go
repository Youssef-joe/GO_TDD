package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord = "Go!"
	write = "write"
	sleep = "sleep"
)

// maybe you can see we're testing too much details in that mocking
// sometimes it would be needed and some other times it would be jus enough
// so, how to determine is it worth or not for detailed testing ?
// and also how much details do i need to test ?
// here's the answer for that : 
// The definition of refactoring is that the code changes but the behaviour stays the same. If you have decided to do some refactoring in theory you should be able to make the commit without any test changes. So when writing a test ask yourself
// Am I testing the behavior I want, or the implementation details?
// If I were to refactor this code, would I have to make lots of changes to the tests?
// Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.
// I feel like if a test is working with more than 3 mocks then it is a red flag - time for a rethink on the design
// Use spies with caution. Spies let you see the insides of the algorithm you are writing which can be very useful but that means a tighter coupling between your test code and the implementation. Be sure you actually care about these details if you're going to spy on them

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}
type DefaultSleeper struct{}

type SpyCountOperations struct {
	Calls []string
}
type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}
type SpyTime struct {
	durationSlept  time.Duration
}
func (s *SpyTime) SetTimeDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountOperations) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, sleep)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1*time.Second)
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i>0; i-- {
		sleeper.Sleep()
	}
	for i := countdownStart; i>0; i-- {
		fmt.Fprint(out, i, " ")

	}
	fmt.Fprint(out, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1*time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}