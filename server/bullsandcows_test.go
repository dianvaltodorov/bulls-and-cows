package main

import "testing"

var isValidTests = []struct {
	b     byte // input
	valid bool // expected result
}{
	{'0', true}, {'1', true}, {'2', true}, {'3', true}, {'4', true},
	{'5', true}, {'6', true}, {'7', true}, {'8', true}, {'9', true},
	{'!', false}, {'a', false}, {'$', false}, {'-', false}, {'%', false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range isValidTests {
		actual := isValid(tt.b)
		if actual != tt.valid {
			t.Errorf("isValid(%v): expected %v, actual %v", tt.b, tt.valid, actual)
		}
	}
}

var beginWithZeroTests = []struct {
	b     []byte // input
	valid bool   // expected result
}{
	{[]byte("0134"), false},
	{[]byte("1234"), true},
	{[]byte(""), false},
}

func TestBeginWithZero(t *testing.T) {
	for _, tt := range isValidTests {
		actual := isValid(tt.b)
		if actual != tt.valid {
			t.Errorf("isValid(%v): expected %v, actual %v", tt.b, tt.valid, actual)
		}
	}
}

var bullsAndCowsNoErrorTests = []struct {
	guess []byte
	bulls int // input
	cows  int // expected result
}{
	{[]byte("1234"), 4, 0},
	{[]byte("4321"), 0, 4},
	{[]byte("1235"), 3, 0},
	{[]byte("1523"), 1, 2},
	{[]byte("1423"), 1, 3},
}

func TestBullsAndCowsNoError(t *testing.T) {
	target := []byte("1234")
	for _, tt := range bullsAndCowsNoErrorTests {
		actualBulls, actualCows, actualErr := bullsAndCows(target, tt.guess)

		if actualBulls != tt.bulls {
			t.Errorf("For bullsAndCows(%v), expected %v bulls, got %v", string(tt.guess), tt.bulls, actualBulls)
		}

		if actualCows != tt.cows {
			t.Errorf("For bullsAndCows(%v), expected %v cows, got %v", string(tt.guess), tt.cows, actualCows)
		}

		if actualErr != nil {
			t.Errorf("For bullsAndCows(%v), Expected error to be nil, got %v", string(tt.guess), actualErr)
		}
	}
}

var bullsAndCowsErrorTests = []struct {
	guess []byte
	err   string
}{
	{[]byte("12434"), "Guess must be 4 characters long"},
	{[]byte("134"), "Guess must be 4 characters long"},
	{[]byte("0321"), "Number can not start with zero"},
	{[]byte("123r"), "Your guess has invalid characters"},
	{[]byte("1123"), "Repeated '1'. No repetition allowed"},
}

func TestBullsAndCowsError(t *testing.T) {
	target := []byte("1234")
	for _, tt := range bullsAndCowsErrorTests {
		actualBulls, actualCows, actualErr := bullsAndCows(target, tt.guess)

		if actualBulls != 0 || actualCows != 0 {
			t.Errorf("For bullsAndCows(%v), expected error state and both bulls and cows to be zero", string(tt.guess))
		}

		if actualErr.Error() != tt.err {
			t.Errorf("For bullsAndCows(%v), Expected error to be %v, got %v ", string(tt.guess), tt.err, actualErr)
		}
	}
}
