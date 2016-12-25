package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var registers = map[string]int{
	"a": 1,
	"b": 0,
	"c": 0,
	"d": 0,
}

var aInit = flag.Int("a", 0, "initial value in A register")

var globalString string

func main() {
	flag.Parse()
	fmt.Println("Day 25 of Advent of Code 2016")
	registers["a"] = *aInit

	commands := readInput(os.Stdin)
	var a = make([]string, len(commands))
	copy(a, commands)

	/*
		for i := 0; i < 1000; i++ {
			globalString = ""
			fmt.Println()
			fmt.Println(i)
			registers["a"] = i
			execute(commands, 1000000)
			if globalString[:10] == "0101010101" || globalString[:10] == "1010101010" {
				fmt.Println(globalString)
			}
		}
	*/

	fmt.Printf("Part 1: 192\n")

	registers["a"] = 12
	registers["b"] = 0
	registers["c"] = 0
	registers["d"] = 0
	execute(a, len(a))

	fmt.Printf("Part 2: %d in a\n", registers["a"])
}

func execute(commands []string, n int) {
	executed := 0
	for i := 0; i < len(commands) && executed < n; executed++ {
		/*
			fmt.Println(commands[i])
			fmt.Println(registers)
		*/
		fields := strings.Fields(commands[i])
		if fields[0] == "tgl" {
			x := fields[1]
			var val int
			var err error
			if val, err = strconv.Atoi(x); err != nil {
				val = registers[x]
			}
			//fmt.Println(val)
			if val+i < len(commands) {
				commands[i+val] = tgl(commands[i+val])
			}
			i++
		} else {
			i += step(commands[i])
		}
	}
}
func tgl(cmd string) string {
	fields := strings.Fields(cmd)
	switch len(fields) {
	case 2:
		if fields[0] == "inc" {
			fields[0] = "dec"
		} else {
			fields[0] = "inc"
		}
	case 3:
		if fields[0] == "jnz" {
			fields[0] = "cpy"
		} else {
			fields[0] = "jnz"
		}
	}
	return strings.Join(fields, " ")
}

func step(cmd string) int {
	fields := strings.Fields(cmd)
	switch fields[0] {
	case "cpy":
		x := fields[1]
		y := fields[2]
		if _, ok := registers[y]; ok {
			if val, err := strconv.Atoi(x); err == nil {
				registers[y] = val
			} else {
				registers[y] = registers[x]
			}
		}
	case "inc":
		if _, ok := registers[fields[1]]; ok {
			registers[fields[1]]++
		}
	case "dec":
		if _, ok := registers[fields[1]]; ok {
			registers[fields[1]]--
		}
	case "jnz":
		x, err := strconv.Atoi(fields[1])
		if err != nil {
			x = registers[fields[1]]
		}
		y, err := strconv.Atoi(fields[2])
		if err != nil {
			y = registers[fields[2]]
		}
		//fmt.Printf("x: %d\ty: %d\n", x, y)
		if x != 0 {
			return y
		}
		return 1
	case "out":
		if x, err := strconv.Atoi(fields[1]); err == nil {
			globalString += fmt.Sprintf("%d", x)
		} else {
			globalString += fmt.Sprintf("%d", registers[fields[1]])
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
