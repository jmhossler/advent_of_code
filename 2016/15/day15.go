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

	input := readInput(f)

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
	for time = 0; !isSuccess(discs, time); time++ {
	}

	fmt.Printf("Part 1: %d\n", time)

	discs = append(discs, disc{11, 0})

	var newTime int
	for newTime = 0; !isSuccess(discs, newTime); newTime++ {
	}

	fmt.Printf("Part 2: %d\n", newTime)
}

func isSuccess(ds []disc, initialTime int) bool {
	for i, disc := range ds {
		if !(rotateDisc(disc, i+initialTime+1) == 0) {
			return false
		}
	}
	return true
}

func rotateDisc(d disc, t int) int {
	return (d.initial + t) % d.positions
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
