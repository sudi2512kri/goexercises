package main

import "fmt"

const (
	spanish       = "Spanish"
	spanishPrefix = "Hola, "
	french        = "French"
	frenchPrefix  = "Bonjour, "
	englishPrefix = "Hello "
)

// private func start with lowerCase
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

// public func's start with UpperCase
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
