package year2021

import (
	"fmt"
	"strings"
)

type day2 struct {
	instruction string
	amount      int
}

func day02ParseInput(input string) []day2 {
	lines := strings.Split(input, "\n")
	items := []day2{}
	for _, v := range lines {
		item := day2{}
		fmt.Sscanf(v, "%s %d", &item.instruction, &item.amount)
		items = append(items, item)
	}

	return items
}

func day02Part1(input string) int {
	items := day02ParseInput(input)
	hor := 0
	depth := 0

	for _, item := range items {
		switch item.instruction {
		case "up":
			depth -= item.amount
		case "down":
			depth += item.amount
		case "forward":
			hor += item.amount
		}
	}
	return hor * depth
}

func day02Part2(input string) int {
	items := day02ParseInput(input)
	hor := 0
	depth := 0
	aim := 0

	for _, item := range items {
		switch item.instruction {
		case "up":
			aim -= item.amount
		case "down":
			aim += item.amount
		case "forward":
			hor += item.amount
			depth += aim * item.amount
		}
	}
	return hor * depth
}
