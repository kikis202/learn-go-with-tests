package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {

		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("counter runs concurently", func(t *testing.T) {
		want := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for range want {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, want)

	})
}

func assertCounter(t testing.TB, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("got %d, want %d", counter.Value(), want)
	}
}
