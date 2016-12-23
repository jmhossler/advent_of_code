package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Day 1 of 2015 Advent of Code")

	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	str := string(data)

	floor := 0
	index := 0
	for i := 0; i < len(str); i++ {
		if floor == -1 && index == 0 {
			index = i
		}
		if str[i] == '(' {
			floor++
		} else if str[i] == ')' {
			floor--
		}
	}

	fmt.Printf("Part 1: Final floor is %v\n", floor)
	fmt.Printf("Part 2: Position of instruction to go to -1 is %v\n", index)
}
