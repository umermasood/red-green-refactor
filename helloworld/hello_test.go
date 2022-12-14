package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when no name is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

		got = Hello("", "French")
		want = "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run(" say hello in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {
		got := Hello("Rene", "French")
		want := "Bonjour, Rene"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in urdu", func(t *testing.T) {
		got := Hello("Abid", "Urdu")
		want := "ہیلو, Abid"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
