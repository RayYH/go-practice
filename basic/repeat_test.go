package basic

import "testing"

func Repeat(character string, repeatCount int) string {
    var repeated string
    for i := 0; i < repeatCount; i++ {
        repeated += character
    }

    return repeated
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}

func TestRepeat(t *testing.T) {
    repeated := Repeat("a", 5)
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}
