package main

import (
	"fmt"
	"teste-capgemini/services"
	"testing"
)

func TestMapMatrixSequence(t *testing.T) {
	listReceived := []string{"BBBBUU", "DDUDDD", "BBBBUU", "DDUDDD", "BBBBUU", "DDUDDD"}
	letter := "B"
	matrix := services.MapSequence(letter, listReceived)
	expectedMatrix := [][]bool{
		{true, true, true, true, false, false},
		{false, false, false, false, false, false},
		{true, true, true, true, false, false},
		{false, false, false, false, false, false},
		{true, true, true, true, false, false},
		{false, false, false, false, false, false},
	}

	expectedResult := fmt.Sprint(expectedMatrix)
	actualResult := fmt.Sprint(matrix)
	if actualResult != expectedResult {
		t.Errorf("Mapping matrix error, want %s, got %s", expectedResult, actualResult)
	}
}
