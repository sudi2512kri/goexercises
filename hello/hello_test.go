package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("start by saying hello to people", func(t *testing.T) {
		got := Hello("juhi", "")
		want := "Hello juhi"

		assertCorrectMessage(t, got, want)
	})

	t.Run("show default hello world if empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Clara", "French")
		want := "Bonjour, Clara"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
