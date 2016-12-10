package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 10 of Advent of Code 2016")

	data := read_data(os.Stdin)
}

func read_lines(f *os.File) []string {
	var data []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
