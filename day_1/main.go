package day1

import (
	"strconv"
)

func SumFirstLastDigits(inputString string) int {
	sum := 0
	firstDigit := 0
	lastDigit := 0
	for _, c := range inputString {
		ch := string(c)
		if ch == "\r" || ch == "\n" {
			sum += firstDigit*10 + lastDigit
			firstDigit = 0
			lastDigit = 0
		}
		num, err := strconv.Atoi(ch)
		if err != nil {
			continue
		}
		if firstDigit == 0 {
			firstDigit = num
		}
		lastDigit = num
	}
	// add remaining (in case of new line is not there in end)
	sum += firstDigit*10 + lastDigit

	return sum
}
