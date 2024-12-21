package main

import "fmt"

const spanish = "Spanish"

const englishWorld = "World"
const englishHelloPrefix = "Hello, "
const spanishWorld = "Mundo"
const spanishHelloPrefix = "Halo, "

const helloSuffix = "!"

func Hello(name, lang string) string {
	if name == "" {
		if lang == spanish {
			name = spanishWorld
		} else {
			name = englishWorld
		}
	}

	if lang == spanish {
		return spanishHelloPrefix + name + helloSuffix
	}

	return englishHelloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("Kristap", ""))
}
