package iteration

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
    iterCount := 5
    result := Iter("a", iterCount)
    expected := "aaaaa"

    if result != expected {
        t.Errorf("expected %q but got %q", expected, result)
    }
}

func TestIterCount(t *testing.T) {
    iterCount := 6
    result := Iter("a", iterCount)
    expected := "aaaaaa"

    if result != expected {
        t.Errorf("expected %q but got %q", expected, result)
    }

}

func BenchmarkIter(b *testing.B) {
    const iterCount = 5

    for i := range b.N {
        Iter("a", iterCount)
        i++
    }
}

func ExampleIter() {
    example := Iter("a", 5)
    fmt.Println(example)
    // Output: aaaaa
}
