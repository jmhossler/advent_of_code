package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 5 solution Advent of Code 2016")

	fp, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		data := scanner.Text()

		fmt.Println(data)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
