package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var registers = map[string]int{
	"a": 0,
	"b": 0,
	"c": 0,
	"d": 0,
}

func main() {
	fmt.Println("Day 12 of Advent of Code 2016")

	commands := readInput(os.Stdin)

	for i := 0; i < len(commands); {
		//fmt.Printf("%d-%s: %v\n", i, commands[i], registers)
		i += execute(commands[i])
	}

	fmt.Printf("Part 1: %d in a\n", registers["a"])

	registers["a"] = 0
	registers["b"] = 0
	registers["d"] = 0
	registers["c"] = 1
	for i := 0; i < len(commands); {
		i += execute(commands[i])
	}

	fmt.Printf("Part 2: %d in a\n", registers["a"])
}

func execute(cmd string) int {
	fields := strings.Fields(cmd)
	switch fields[0] {
	case "cpy":
		x := fields[1]
		y := fields[2]
		if val, err := strconv.Atoi(x); err == nil {
			registers[y] = val
		} else {
			registers[y] = registers[x]
		}
	case "inc":
		registers[fields[1]]++
	case "dec":
		registers[fields[1]]--
	case "jnz":
		x, err := strconv.Atoi(fields[1])
		if err != nil {
			x = registers[fields[1]]
		}
		y, err := strconv.Atoi(fields[2])
		if err != nil {
			y = registers[fields[1]]
		}
		if x != 0 {
			return y
		}
	}
	return 1
}

func readInput(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
