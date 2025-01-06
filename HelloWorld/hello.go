package main

import "fmt"

const (
    spanish = "Spanish"
    french = "French"
    EnglishHelloPrefix = "Hello, "
    SpanishHelloPrefix = "Hola, "
    FrenchHelloPrefix = "Bonjour, "
)
func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language)+name
}

func greetingPrefix(language string) (prefix string) {
    switch language {
    case spanish:
        prefix = SpanishHelloPrefix
    case french:
        prefix = FrenchHelloPrefix
    default:
        prefix = EnglishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("Hox", "English"))
}
