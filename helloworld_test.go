package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	msg := Greeting()
	if msg != "Hello world\n" {
		t.Fatalf("failed test %#v", msg)
	}
}
