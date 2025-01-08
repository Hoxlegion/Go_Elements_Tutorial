package main

import "testing"

func TestPerimeter(t *testing.T) {
    result := Perimeter(10.0, 10.0)
    want := 40.0

    if result != want {
        t.Errorf("got %.2f want %.2f", result, want)
    }
}
