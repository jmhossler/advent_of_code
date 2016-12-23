package main

import (
	"flag"
	"fmt"
)

var numElves = flag.Int("n", 3017957, "Number of elves")

func main() {
	flag.Parse()
	//fmt.Println("Day 19 of Advent of Code 2016")

	elf1 := getElf(*numElves)

	elf2 := getElfPt2(*numElves)
	fmt.Printf("%d-%d-%d\n", *numElves, elf1, elf2)
}

func getElfPt2(numElves int) int {
	elf := 1
	for i := 1; i < numElves; i++ {
		elf = elf%i + 1
		if elf > (i+1)/2 {
			elf++
		}
	}
	return elf
}

func getElf(numElves int) int {
	binary := fmt.Sprintf("%b", numElves)
	sol := 1
	for i := 1; i < len(binary); i++ {
		if binary[i] == '1' {
			sol += 1 << uint(len(binary)-i)
		}
	}

	return sol
}
