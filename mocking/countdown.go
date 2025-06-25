package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const startingNumber = 3
const finalWord = "Go!"

func Countdown(writer io.Writer) {
	for i := startingNumber; i > 0; i-- {
		fmt.Fprintln(writer, fmt.Sprintf("%d,", i))
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
