package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 6 of Advent of Code 2016")
	f, _ := os.Open("input")

	m := [8]map[byte]int{}
	for i := 0; i < 8; i++ { // init each index of m
		m[i] = make(map[byte]int)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			m[i][line[i]]++
		}
	}

	var text []byte
	var pt2Text []byte
	for i := 0; i < 8; i++ {
		text = append(text, getExtremeKey(m[i], func(x, y int) bool { return x > y }, -1))
		pt2Text = append(pt2Text, getExtremeKey(m[i], func(x, y int) bool { return x < y }, 1))
	}

	fmt.Printf("Part 1: %s\n", text)
	fmt.Printf("Part 2: %s\n", pt2Text)
}

func getExtremeKey(m map[byte]int, f func(x, y int) bool, parity int) byte {
	extreme := parity * 1000
	var key byte
	for k, v := range m {
		if f(v, extreme) {
			extreme = v
			key = k
		}
	}
	return key
}
