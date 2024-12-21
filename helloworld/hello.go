package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishWorld       = "World"
	englishHelloPrefix = "Hello, "
	spanishWorld       = "Mundo"
	spanishHelloPrefix = "Halo, "
	frenchWorld        = "Monde"
	frenchHelloPrefix  = "Bonjour, "

	helloSuffix = "!"
)

func Hello(name, lang string) string {
	if name == "" {
		worldTranslation(&name, lang)
	}

	return greetingPrefix(lang) + name + helloSuffix
}

func worldTranslation(name *string, lang string) {
	switch lang {
	case spanish:
		*name = spanishWorld
	case french:
		*name = frenchWorld
	default:
		*name = englishWorld
	}
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Kristap", ""))
}
