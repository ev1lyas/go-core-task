package main

import "testing"

func TestRandomNumberGenerator(t *testing.T) {
	ch := RandomNumberGenerator()
	val1 := <-ch
	val2 := <-ch
	if val1 == val2 {
		t.Errorf("Expected different random values, but got %v and %v", val1, val2)
	}
}
