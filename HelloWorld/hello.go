package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }
    if language == spanish {
        return SpanishHelloPrefix+name
    } else if language == french {
        return FrenchHelloPrefix+name
    }
    return EnglishHelloPrefix+name
}

func main() {
    fmt.Println(Hello("Hox", "English"))
}
