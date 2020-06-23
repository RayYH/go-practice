package basic

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Ray", englishHelloPrefix)
		want := "Hello, Ray"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", englishHelloPrefix)
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in chinese", func(t *testing.T) {
		got := Hello("Ray", "Chinese")
		want := "你好, Ray"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Ray", "Spanish")
		want := "Hola, Ray"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Ray", "French")
		want := "Bonjour, Ray"
		assertCorrectMessage(t, got, want)
	})
}
