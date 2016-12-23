package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rnge struct {
	low, high uint
}

func main() {
	fmt.Println("Day 20 of Advent of Code 2016")

	input := readInput(os.Stdin)

	input = mergeSort(input)
	//input = groupRanges(input)
	lowest := getLowestSlot(input, 0)
	fmt.Printf("Part 1: %d\n", lowest)
	lenSlots := getLengthSlots(input)
	fmt.Printf("Part 2: %d\n", lenSlots)
}

func readInput(f *os.File) []rnge {
	scanner := bufio.NewScanner(f)
	data := []rnge{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "-")
		low, _ := strconv.ParseInt(fields[0], 10, 64)
		high, _ := strconv.ParseInt(fields[1], 10, 64)
		data = append(data, rnge{low: uint(low), high: uint(high)})
	}
	return data
}

func getLowestSlot(lines []rnge, base uint) uint {
	lowest := base
	for _, r := range lines {
		if inRange(r.low, r.high, lowest) {
			lowest = r.high + 1
		}
	}
	return lowest
}

func groupRanges(lines []rnge) []rnge {
	newRange := []rnge{}
	for i := 0; i < len(lines); i++ {
		high := lines[i].high
		var j int
		for j = i + 1; j < len(lines) && inRange(lines[i].low, high+1, lines[j].low); j++ {
			high = lines[j].high
		}
		newRange = append(newRange, rnge{lines[i].low, high})
		i = j - 1
	}
	return newRange
}

func getLengthSlots(lines []rnge) uint {
	length := lines[0].low
	lowest := lines[0].low
	high := lines[0].high
	for i := 0; i < len(lines); i++ {
		if lines[i].low > high+1 {
			length += high - lowest + 1
			lowest = lines[i].low
			high = lines[i].high
		} else {
			high = max(high, lines[i].high)
		}
	}

	return 4294967296 - length - (high - lowest + 1)
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func shrinkRange(lines []rnge, low uint) []rnge {
	if len(lines) > 0 && lines[0].low > low {
		return lines[1:]
	}
	return lines
}

func mergeSort(r []rnge) []rnge {
	if len(r) == 1 {
		return r
	}
	newRange := make([]rnge, len(r))
	left := mergeSort(r[:len(r)/2])
	right := mergeSort(r[len(r)/2:])
	for i := 0; i < len(r); i++ {
		if len(left) != 0 && (len(right) == 0 || left[0].low < right[0].low) {
			newRange[i] = left[0]
			left = left[1:]
		} else if len(left) == 0 || left[0].low > right[0].low {
			newRange[i] = right[0]
			right = right[1:]
		} else if left[0].high < right[0].high {
			newRange[i] = left[0]
			left = left[1:]
		} else {
			newRange[i] = right[0]
			right = right[1:]
		}
	}
	return newRange
}

func inRange(low, high, val uint) bool {
	return val >= low && val <= high
}
