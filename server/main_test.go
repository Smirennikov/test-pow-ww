package main

import (
	"fmt"
	"testing"
)

func TestGenerateChallenge(t *testing.T) {
	challenge := generateChallenge()
	if challenge == "" {
		t.Errorf("Generated challenge doesn't have the correct prefix")
	}
}

func TestVerifyProofOfWork(t *testing.T) {
	validSolution := "2194970410229553659659240"
	invalidSolution := "123456"

	fmt.Println("test", verifyProofOfWork(validSolution))
	if !verifyProofOfWork(validSolution) {
		t.Errorf("Valid solution was rejected")
	}
	if verifyProofOfWork(invalidSolution) {
		t.Errorf("Invalid solution was accepted")
	}
}

func TestGetRandomQuote(t *testing.T) {

	gotQuotes := make(map[string]int, len(quotes))
	for _, q := range quotes {
		gotQuotes[q] = 0
	}

	var prevQuote string
	for i := 2; i > 0; i-- {
		quote := getRandomQuote()

		if prevQuote == quote {
			t.Errorf("Generated rand quote give same result")
		}
	}
}
