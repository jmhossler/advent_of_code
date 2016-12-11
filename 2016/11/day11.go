package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 11 of Advent of Code 2016")
	lines := get_lines(os.Stdin)
}

func get_lines(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
