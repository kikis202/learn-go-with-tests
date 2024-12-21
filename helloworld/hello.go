package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloSuffix = "!"

func Hello(name string) string {
	return englishHelloPrefix + name + englishHelloSuffix
}

func main() {
	fmt.Println(Hello("Kristap"))
}
