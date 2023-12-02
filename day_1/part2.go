package day1

import (
	"slices"
	"strconv"
	"unicode"
)

func SumFirstLastDigitsAndWords(inputString string) int {
	digitWords := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0
	firstDigit := 0
	lastDigit := 0
	charArr := []rune(inputString)
	for i, c := range inputString {
		ch := string(c)
		if ch == "\r" || ch == "\n" {
			sum += firstDigit*10 + lastDigit
			firstDigit = 0
			lastDigit = 0
		}
		num, err := strconv.Atoi(ch)
		if err != nil {
			word := ""
			j := i
			for j < len(charArr) {
				wordChar := string(charArr[j])
				if wordChar == "\r" || wordChar == "\n" || unicode.IsDigit(charArr[j]) {
					break
				}
				word += wordChar
				indx := slices.Index(digitWords, word)
				if indx != -1 {
					if firstDigit == 0 {
						firstDigit = indx + 1
					}
					lastDigit = indx + 1
					break
				}
				j++
			}
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
