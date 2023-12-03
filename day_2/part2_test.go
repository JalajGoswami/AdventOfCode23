package day2

import (
	"os"
	"testing"
)

func TestSumPowerSetOfCube(t *testing.T) {
	result := SumPowerSetOfCube(testInput)
	if result != 2286 {
		t.Fatalf("Test SumPowerSetOfCube() failed expected %d got %d", 2286, result)
	}
	data, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	result = SumPowerSetOfCube(string(data))
	if result != 78111 {
		t.Fatalf("Test SumPowerSetOfCube() failed expected %d got %d", 78111, result)
	}
}
