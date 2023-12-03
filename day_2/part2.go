package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func SumPowerSetOfCube(inputString string) int {
	sum := 0
	// CRLF to LF
	inputString = strings.TrimSpace(strings.ReplaceAll(inputString, "\r\n", "\n"))
	lines := strings.Split(inputString, "\n")

	for _, line := range lines {
		gameInfo := strings.Split(line, ":")
		results := strings.Split(gameInfo[1], ";")
		maxCubeCount := map[string]int{}
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
				if count > maxCubeCount[color] {
					maxCubeCount[color] = count
				}
			}
		}
		sum += maxCubeCount["red"] * maxCubeCount["green"] * maxCubeCount["blue"]
	}
	return sum
}
