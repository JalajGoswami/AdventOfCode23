package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func SumFeasibleGameIds(inputString string) int {
	maxCubeCount := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
	// CRLF to LF
	inputString = strings.TrimSpace(strings.ReplaceAll(inputString, "\r\n", "\n"))
	lines := strings.Split(inputString, "\n")

lineLoop:
	for _, line := range lines {
		gameInfo := strings.Split(line, ":")
		gameId, err := strconv.Atoi(strings.Split(gameInfo[0], " ")[1])
		if err != nil {
			panic(fmt.Errorf("Game id not a number! Error=%w", err))
		}
		results := strings.Split(gameInfo[1], ";")
		for _, cubeCounts := range results {
			colorCounts := strings.Split(cubeCounts, ",")
			for _, colorCount := range colorCounts {
				colorCount = strings.TrimSpace(colorCount)
				colorCountArr := strings.Split(colorCount, " ")
				color := colorCountArr[1]
				count, err := strconv.Atoi(colorCountArr[0])
				if err != nil {
					panic(fmt.Errorf("Cube count not a number! Error=%w", err))
				}
				maxCount, exist := maxCubeCount[color]
				if !exist {
					panic(fmt.Errorf("Invalid cube color! Error=%w", err))
				}
				if count > maxCount {
					continue lineLoop
				}
			}
		}
		sum += gameId
	}
	return sum
}
