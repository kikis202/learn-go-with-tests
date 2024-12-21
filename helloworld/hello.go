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
	prefix := englishHelloPrefix
	switch lang {
	case spanish:
		if name == "" {
			name = spanishWorld
		}
		prefix = spanishHelloPrefix
	case french:
		if name == "" {
			name = frenchWorld
		}
		prefix = frenchHelloPrefix
	default:
		if name == "" {
			name = englishWorld
		}
	}

	return prefix + name + helloSuffix
}

func main() {
	fmt.Println(Hello("Kristap", ""))
}
