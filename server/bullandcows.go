package main

import (
	"bytes"
	"errors"
	"fmt"
)

var alphabet = []byte("0123456789")

const allowedLength = 4

func validLength(guess []byte) bool {
	return len(guess) != allowedLength
}

func beginWithZero(guess []byte) bool {
	return len(guess) > 0 && guess[0] == '0'
}

func isValid(c byte) bool {
	return bytes.IndexByte(alphabet, c) != -1
}

func bullsAndCows(target []byte, guess []byte) (int, int, error) {
	var bulls, cows int

	if validLength(guess) {
		return 0, 0, fmt.Errorf("Guess must be %d characters long", allowedLength)
	}

	if beginWithZero(guess) {
		return 0, 0, fmt.Errorf("Number can not start with zero")
	}

	for idx, c := range guess {

		if !isValid(c) {
			return 0, 0, errors.New("Your guess has invalid characters")
		}

		if bytes.IndexByte(guess[:idx], c) >= 0 {
			return 0, 0, fmt.Errorf("Repeated '%c'. No repetition allowed", c)
		}

		pos := bytes.IndexByte(target, c)

		if pos == idx {
			bulls = bulls + 1
		} else {
			if pos >= 0 {
				cows = cows + 1
			}
		}
	}
	return bulls, cows, nil
}
