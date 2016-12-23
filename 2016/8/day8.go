package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var screen = [6][50]bool{}

func main() {
	fmt.Println("Day 8 of Advent of Code 2016")
	f, err := os.Create("screen.pgm")
	check(err)

	information := readInput("input")
	for _, line := range information {
		runCommand(line)
	}
	writeImage(f)

	fmt.Printf("%d pixels on\n", countOn())
}

func writeImage(f *os.File) {
	w := bufio.NewWriter(f)
	w.WriteString("P2\n50 6\n255\n")
	for _, row := range screen {
		for _, val := range row {
			if val {
				w.WriteString("0 ")
			} else {
				w.WriteString("255 ")
			}
		}
		w.WriteString("\n")
	}
	w.Flush()
}

func runCommand(cmd string) {
	fields := strings.Fields(cmd)
	switch fields[0] {
	case "rect":
		info := strings.Split(fields[1], "x")
		a, _ := strconv.Atoi(info[0])
		b, _ := strconv.Atoi(info[1])
		callRect(a, b)
	case "rotate":
		info := strings.Split(fields[2], "=")
		a, _ := strconv.Atoi(info[1])
		b, _ := strconv.Atoi(fields[4])
		if fields[1] == "row" {
			rotateRow(a, b)
		} else {
			rotateCol(a, b)
		}
	}
}

func countOn() int {
	var count int
	for _, row := range screen {
		for _, pix := range row {
			if pix {
				count++
			}
		}
	}
	return count
}

func callRect(a, b int) {
	for i := 0; i < b && i < len(screen); i++ {
		for j := 0; j < a && j < len(screen[i]); j++ {
			screen[i][j] = true
		}
	}
}

func rotateRow(a, b int) {
	newCol := make([]bool, 50)
	for i := 0; i < 50; i++ {
		newCol[(i+b)%50] = screen[a][i]
	}
	for i := 0; i < 50; i++ {
		screen[a][i] = newCol[i]
	}
}

func rotateCol(a, b int) {
	newCol := make([]bool, 6)
	for i := 0; i < 6; i++ {
		newCol[(i+b)%6] = screen[i][a]
	}
	for i := 0; i < 6; i++ {
		screen[i][a] = newCol[i]
	}
}

func shiftRow(a int) {
	a = a % len(screen[a])
	temp := screen[a][0]
	for i := 1; i < len(screen[a]); i++ {
		temp, screen[a][i] = screen[a][i], temp
	}
	screen[a][0] = temp
}

func shiftCol(a int) {
	a = a % len(screen)
	temp := screen[0][a]
	for i := 1; i < len(screen); i++ {
		temp, screen[i][a] = screen[i][a], temp
	}
	screen[0][a] = temp
}

func readInput(filename string) []string {
	fp, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(fp)

	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
