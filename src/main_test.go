package main

import "testing"

func TestGreeting(t *testing.T) {
	if greeting("a") != "<b>a</b>" {
		t.Error("Tag error.")
	}
}

