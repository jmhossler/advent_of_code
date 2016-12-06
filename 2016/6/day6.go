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

	var m []map[byte]int

	for i := 0; i < 8; i++ {
		new_map := make(map[byte]int)
		m = append(m, new_map)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			m[i][line[i]]++
		}
	}

	var text []byte
	var pt2_text []byte
	for i := 0; i < 8; i++ {
		text = append(text, get_highest(m[i]))
		pt2_text = append(pt2_text, get_lowest(m[i]))
	}

	fmt.Printf("Part 1: %s\n", text)
	fmt.Printf("Part 2: %s\n", pt2_text)
}

func get_lowest(m map[byte]int) byte {
	lowest := 10000
	var store byte
	for k, v := range m {
		if v < lowest {
			lowest = v
			store = k
		}
	}
	return store
}

func get_highest(m map[byte]int) byte {
	var highest int
	var store byte
	for k, v := range m {
		if v > highest {
			highest = v
			store = k
		}
	}
	return store
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
