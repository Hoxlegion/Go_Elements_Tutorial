package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hola to the panish", func(t *testing.T) {
        var name, language string = "Hox", "Spanish"

        got := Hello(name, language)
        want := "Hola, "+name
        assertCorrectMessage(t, got, want)
    })
    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        var name, language string = "", ""
        got := Hello(name, language)
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })
    t.Run("say 'Bonjour, World' French is supplied as language", func(t *testing.T) {
        var name, language string = "Franciare", "French"
        got := Hello(name, language)
        want := "Bonjour, Franciare"
        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

