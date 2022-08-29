package main

import "fmt"

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "Spanish" {
		return fmt.Sprintf("Hola, %v!", name)
	}
	return fmt.Sprintf("Hello, %v!", name)
}

func main() {}
