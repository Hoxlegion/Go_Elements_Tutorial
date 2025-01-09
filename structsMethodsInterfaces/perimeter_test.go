package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
    rectangle := Rectangle{10.0, 10.0}
    result := Perimeter(rectangle)
    want := 40.0

    if result != want {
        t.Errorf("got %.2f want %.2f", result, want)
    }
}

func TestArea(t *testing.T) {

    areaTest := []struct {
        name string
        shape Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
        {name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{5, 6}, hasArea: 15},
    }

    for _, tt := range areaTest {
        t.Run("Shape testing", func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g hasArea %g", tt.name, got, tt.hasArea)
            }
        })
    }
//    checkArea := func(t testing.TB, shape Shape, want float64) {
//        t.Helper()
//        got := shape.Area()
//        if got != want {
//            t.Errorf("got %g want %g", got, want)
//        }
//    }
//
//    t.Run("rectangles", func (t *testing.T) {
//        rectangle := Rectangle{12.0, 6.0}
//        checkArea(t, rectangle, 72.0)
//    })
//
//    t.Run("circles", func (t *testing.T) {
//        circle := Circle{10}
//        checkArea(t, circle, 314.1592653589793)
//    })
//
//    t.Run("triangle", func (t *testing.T) {
//        triangle := Triangle{5, 6}
//        checkArea(t, triangle, 15)
//    })
}
