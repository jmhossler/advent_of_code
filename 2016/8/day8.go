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

	information := read_input("input")
	for _, line := range information {
		run_command(line)
	}
	write_image(f)

	fmt.Printf("%d pixels on\n", count_on())
}

func write_image(f *os.File) {
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

func run_command(cmd string) {
	fields := strings.Fields(cmd)
	switch fields[0] {
	case "rect":
		info := strings.Split(fields[1], "x")
		a, _ := strconv.Atoi(info[0])
		b, _ := strconv.Atoi(info[1])
		call_rect(a, b)
	case "rotate":
		info := strings.Split(fields[2], "=")
		a, _ := strconv.Atoi(info[1])
		b, _ := strconv.Atoi(fields[4])
		if fields[1] == "row" {
			rotate_row(a, b)
		} else {
			rotate_col(a, b)
		}
	}
}

func count_on() int {
	var count int
	for _, row := range screen {
		for _, pix := range row {
			if pix {
				count += 1
			}
		}
	}
	return count
}

func call_rect(a, b int) {
	for i := 0; i < b && i < len(screen); i++ {
		for j := 0; j < a && j < len(screen[i]); j++ {
			screen[i][j] = true
		}
	}
}

func rotate_row(a, b int) {
	new_col := make([]bool, 50)
	for i := 0; i < 50; i++ {
		new_col[(i+b)%50] = screen[a][i]
	}
	for i := 0; i < 50; i++ {
		screen[a][i] = new_col[i]
	}
}

func rotate_col(a, b int) {
	new_col := make([]bool, 6)
	for i := 0; i < 6; i++ {
		new_col[(i+b)%6] = screen[i][a]
	}
	for i := 0; i < 6; i++ {
		screen[i][a] = new_col[i]
	}
}

func shift_row(a int) {
	a = a % len(screen[a])
	temp := screen[a][0]
	for i := 1; i < len(screen[a]); i++ {
		temp, screen[a][i] = screen[a][i], temp
	}
	screen[a][0] = temp
}

func shift_col(a int) {
	a = a % len(screen)
	temp := screen[0][a]
	for i := 1; i < len(screen); i++ {
		temp, screen[i][a] = screen[i][a], temp
	}
	screen[0][a] = temp
}

func read_input(filename string) []string {
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
