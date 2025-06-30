package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureRespTime(a)
	durationB := measureRespTime(b)

	if durationA < durationB {
		return a
	}
	return b
}

func measureRespTime(url string) (duration time.Duration) {
	start := time.Now()
	http.Get(url)
	return time.Since(start)

}
