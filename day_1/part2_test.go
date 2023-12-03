package day1

import (
	"os"
	"testing"
)

func TestSumFirstLastDigitsAndWords(t *testing.T) {
	result := SumFirstLastDigitsAndWords(testInput)
	if result != 281 {
		t.Fatalf("Test SumFirstLastDigitsAndWords() failed expected %d got %d", 281, result)
	}
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	result = SumFirstLastDigitsAndWords(string(data))
	if result != 55614 {
		t.Fatalf("Test SumFirstLastDigitsAndWords() failed expected %d got %d", 55614, result)
	}
}
