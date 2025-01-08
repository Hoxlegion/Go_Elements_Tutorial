package main

import "testing"

type Rectangle struct {
    Width float64
    Height float64
}

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    result := Perimeter(rectangle)
    want := 40.0

    if result != want {
        t.Errorf("got %.2f want %.2f", result, want)
    }
}

func TestArea(t *testing.T) {
    rectangle := Rectangle{12.0, 6.0}
    result := Area(rectangle)
    want := 72.0

    if result != want {
        t.Errorf("got %v want %v", result, want)
    }
}
