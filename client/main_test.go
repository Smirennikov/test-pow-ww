package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

func TestSolveProofOfWork(t *testing.T) {
	challenge := "123456"
	solution := solveProofOfWork(challenge)
	if !verifyProofOfWork(solution) {
		t.Errorf("Generated solution is not valid")
	}
}

func verifyProofOfWork(solution string) bool {
	hash := sha256.Sum256([]byte(solution))
	hashString := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hashString, strings.Repeat("0", difficulty))
}
