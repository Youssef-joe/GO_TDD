package main

import "fmt"

const (
	german        = "German"
	spanish       = "Spanish"
	englishPrefix = "Hello, "
	germanPrefix  = "Hallo, "
	spanishPrefix = "Hola, "
)

func Hello(name string, language string) string {
	return getLang(language) + name
}

// that's in lower-case cuz it's minimal/helper function
func getLang(language string) (prefix string) {
	switch language {
	case german:
		prefix = germanPrefix

	case spanish:
		prefix = spanishPrefix

	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "German"))
}
