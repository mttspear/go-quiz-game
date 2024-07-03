package main

import (
	"testing"
)

// TestCorrectScoreAnswer that correct answers return true
func TestCorrectScoreAnswer(t *testing.T) {
	scoreAnswer := "5"
	questionAnswer := "5"
	result := ScoreAnswer(scoreAnswer, questionAnswer)
	if !result {
		t.Fatalf(`Incorrect answer being returned as correct`)
	}
}

// TestIncorrectScoreAnswer that correct answers return false
func TestIncorrectScoreAnswer(t *testing.T) {
	scoreAnswer := "5"
	questionAnswer := "6"
	result := ScoreAnswer(scoreAnswer, questionAnswer)
	if result {
		t.Fatalf(`Incorrect answer being returned as correct`)
	}
}
