package main

import "testing"

func TestChoseNumber(t *testing.T) {
	num := genNumber()
	if len(num) != AllowedLength {
		t.Errorf("Expected length of chosen number to be %d", AllowedLength)
	}
}
