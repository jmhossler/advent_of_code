package main

import (
	"flag"
	"fmt"
)

var numElves = flag.Int("n", 3017957, "Number of elves")

func main() {
	flag.Parse()
	//fmt.Println("Day 19 of Advent of Code 2016")

	var elf int
	//elf = getElf(*numElves)
	//fmt.Printf("Part 1: %d\n", elf)

	elf = getElfPt2(*numElves)
	fmt.Printf("Part 2: %d elves - %d final\n", *numElves, elf)
}

func getElfPt2(numElves int) int {
	elf := 1
	for i := 1; i < numElves; i++ {
		elf = elf%i + 1
		if elf > (i+1)/2 {
			elf += 1
		}
	}
	return elf
}

func getElf(numElves int) int {
	elves := make([]int, numElves)
	for i := 0; i < numElves; i++ {
		elves[i] = i + 1
	}

	for i := 0; len(elves) > 1; i = (i + 1) % (len(elves) + 1) {
		//fmt.Println(elves)
		//fmt.Println(i)
		if i+1 >= len(elves) {
			elves = elves[1:]
		} else if i+2 >= len(elves) {
			elves = elves[:len(elves)-1]
		} else {
			elves = append(elves[:(i+1)%(len(elves))], elves[(i+2)%(len(elves)):]...)
		}
	}
	return elves[0]
}
