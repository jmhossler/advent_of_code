package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type disc struct {
	positions, initial int
}

func main() {
	fmt.Println("Day 15 of Advent of Code 2016")
	f, err := os.Open("input")
	check(err)

	input := read_input(f)

	var discs []disc
	for _, line := range input {
		fields := strings.Fields(line)
		positions, err := strconv.Atoi(fields[3])
		check(err)
		initial, err := strconv.Atoi(fields[11][:len(fields[11])-1])
		check(err)
		discs = append(discs, disc{positions, initial})
	}

	var time int
	for time = 0; !is_success(discs, time); time++ {
	}

	fmt.Printf("Part 1: %d\n", time)

	discs = append(discs, disc{11, 0})

	var new_time int
	for new_time = 0; !is_success(discs, new_time); new_time++ {
	}

	fmt.Printf("Part 2: %d\n", new_time)
}

func is_success(ds []disc, initial_time int) bool {
	for i, disc := range ds {
		if !(rotate_disc(disc, i+initial_time+1) == 0) {
			return false
		}
	}
	return true
}

func rotate_disc(d disc, t int) int {
	return (d.initial + t) % d.positions
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_input(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
