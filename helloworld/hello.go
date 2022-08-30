package main

const frenchHelloPrefix = "Bonjour"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const urduHelloPrefix = "ہیلو"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Urdu":
		prefix = urduHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
