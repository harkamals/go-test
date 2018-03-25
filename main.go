package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const helloPrefixEnglish = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpanish
	case french:
		prefix = helloPrefixFrench
	default:
		prefix = helloPrefixEnglish
	}
	return prefix
}

//Hello is exported
func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("World", spanish))
}
