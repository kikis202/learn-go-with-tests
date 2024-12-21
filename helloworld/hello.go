package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishWorld = "World"
const englishHelloPrefix = "Hello, "
const spanishWorld = "Mundo"
const spanishHelloPrefix = "Halo, "
const frenchWorld = "Monde"
const frenchHelloPrefix = "Bonjour, "

const helloSuffix = "!"

func Hello(name, lang string) string {
	if name == "" {
		if lang == spanish {
			name = spanishWorld
		} else if lang == french {
			name = frenchWorld
		} else {
			name = englishWorld
		}
	}

	if lang == spanish {
		return spanishHelloPrefix + name + helloSuffix
	}
	if lang == french {
		return frenchHelloPrefix + name + helloSuffix
	}

	return englishHelloPrefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("Kristap", ""))
}
