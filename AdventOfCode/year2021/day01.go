package year2021

import (
	"strconv"
	"strings"
)

func day01ParseInput(input string) map[int]int {
	lines := strings.Split(input, "\n")
	numLines := map[int]int{}
	for k, v := range lines {
		i, _ := strconv.Atoi(v)
		numLines[k] = i
	}

	return numLines
}

func day01Part1(input string) int {
	lines := day01ParseInput(input)

	answer := 0
	for k, v := range lines {
		if k == 0 {
			continue
		}
		if v > lines[k-1] {
			answer++
		}
	}
	return answer
}

func day01Part2(input string) int {
	lines := day01ParseInput(input)

	answer := 0
	for k, _ := range lines {
		if lines[k]+lines[k+1]+lines[k+2] < lines[k+1]+lines[k+2]+lines[k+3] {
			answer++
		}
	}
	return answer
}
