package basic

import (
    "fmt"
    "testing"
)

func Add(x, y int) int {
    return x + y
}

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4
    if sum != expected {
        t.Errorf("expect '%d' but got '%d'", expected, sum)
    }
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
