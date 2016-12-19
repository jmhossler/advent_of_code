package main

import (
	"flag"
	"fmt"
)

var numElves = flag.Int("n", 3017957, "Number of elves")

func main() {
	flag.Parse()
	fmt.Println("Day 19 of Advent of Code 2016")

	var elf int
	elf = getElf(*numElves)
	fmt.Printf("Part 1: %d\n", elf)

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
	pow := 1
	for pow*2 < numElves {
		pow = pow * 2
	}
	return 2*(numElves-pow) + 1
}
