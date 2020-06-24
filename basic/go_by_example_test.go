package basic

import "testing"

// https://gobyexample.com/

func TestIntMin(t *testing.T) {
	got := IntMin(2, -2)
	expected := -2

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
