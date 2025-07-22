package main

import "fmt"

const englishPrefix = "Hello, "
const japanesePrefix = "こんにちは, "
const japanese = "Japanese"

func Hello(language string, name string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case japanese:
		return japanesePrefix + name
	default:
		return englishPrefix + name
	}
}

func main() {
	fmt.Println(Hello("English", "world"))
}
