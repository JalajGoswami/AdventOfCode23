package day1

import (
	"os"
	"testing"
)

const testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestSumFirstLastDigits(t *testing.T) {
	result := SumFirstLastDigits(testInput)
	if result != 142 {
		t.Fatalf("Test SumFirstLastDigits() failed expected %d got %d", 142, result)
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
