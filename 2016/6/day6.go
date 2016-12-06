package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 6 of Advent of Code 2016")

	f, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
