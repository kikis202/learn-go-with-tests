package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloSuffix = "!"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name + englishHelloSuffix
}

func main() {
	fmt.Println(Hello("Kristap"))
}
