package main 

import "testing"

func TestWelcome(t *testing.T) {
	want := "Welcome"
	if got := Welcome(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
