package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var rows = flag.Int("n", 40, "rows")

func main() {
	flag.Parse()
	fmt.Println("Day 18 of Advent of Code 2016")

	input := readInput(os.Stdin)

	floor := extendRow(input, *rows)
	var count int
	for _, s := range floor {
		for _, c := range s {
			if c == '.' {
				count += 1
			}
		}
	}

	fmt.Printf("%d safe tiles\n", count)
}

func extendRow(s string, length int) []string {
	floor := []string{s}
	for i := 0; i < length-1; i++ {
		nextFloor := getNextRow(floor[i])
		floor = append(floor, nextFloor)
	}
	return floor
}

func getNextRow(s string) string {
	newString := []rune{}

	for i := 0; i < len(s); i++ {
		left := isTrap(i-1, s)
		center := isTrap(i, s)
		right := isTrap(i+1, s)
		if shouldTrap(left, center, right) {
			newString = append(newString, '^')
		} else {
			newString = append(newString, '.')
		}
	}

	return string(newString)
}

func shouldTrap(left, center, right bool) bool {
	if left && center && !right {
		return true
	} else if !left && center && right {
		return true
	} else if left && !center && !right {
		return true
	} else if !left && !center && right {
		return true
	} else {
		return false
	}
}

func isTrap(index int, s string) bool {
	if index < 0 || index >= len(s) {
		return false
	} else {
		return s[index] == '^'
	}
}

func readInput(f *os.File) string {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}
