package main

import (
	"os"
	"time"

	clockface "github.com/kikis202/learn-go-with-tests/learnGoWithTests/math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
