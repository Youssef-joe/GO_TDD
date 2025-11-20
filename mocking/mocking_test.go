package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("testing the fallback result", func(t *testing.T) {
		buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := "3 2 1 Go!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 but got %d", spySleeper.Calls)
	}
	})

	t.Run("testing sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			sleep,
			sleep,
			sleep,
			sleep,
			sleep,
			sleep,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5*time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetTimeDurationSlept}
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}