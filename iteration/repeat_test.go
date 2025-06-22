package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Print(repeated)
	// Output: aaa
}
