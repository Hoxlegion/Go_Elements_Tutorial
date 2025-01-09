package main

type Triangle struct {
    Width float64
    Length float64
}

func (t Triangle) Area() float64 {
    return (t.Width * t.Length) / 2
}
