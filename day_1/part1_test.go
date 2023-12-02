package day1

import (
	"os"
	"testing"
)

func TestSumFirstLastDigits(t *testing.T) {
	result := SumFirstLastDigits(testInput)
	if result != 281 {
		t.Fatalf("Test SumFirstLastDigits() failed expected %d got %d", 281, result)
	}
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	result = SumFirstLastDigits(string(data))
	if result != 55488 {
		t.Fatalf("Test SumFirstLastDigits() failed expected %d got %d", 55488, result)
	}
}
