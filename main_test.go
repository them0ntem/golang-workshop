package main

import "testing"

func TestHelloValidName(t *testing.T) {
	name := "Shreyas"
	want := "Hello Shreyas!"
	msg, err := Hello(name)
	if msg != want || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloInvalidName(t *testing.T) {
	name := ""
	_, err := Hello(name)
	if err == nil {
		t.Fatalf(`Hello("") returned no error`)
	}
}
