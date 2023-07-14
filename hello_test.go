package main

import "testing"

func TestHello(t *testing.T) {
	//相当于Java中的  assertEquals
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
