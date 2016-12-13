package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Floor []Item

type Item struct {
	name        string
	description string
}

type State struct {
	floors          []Floor
	distance, floor int
}

func main() {
	fmt.Println("Day 11 of Advent of Code 2016")
	var floors = read_input(os.Stdin)

	min := min_path(floors)

	fmt.Printf("Part 1: %d\n", min)
}

func read_input(f *os.File) []Floor {
	scanner := bufio.NewScanner(f)
	var floors []Floor
	for scanner.Scan() {
		line := scanner.Text()
		var floor Floor
		info := strings.Fields(line)
		for i := range info {
			if info[i] == "a" {
				if info[i+2][:len(info[i+2])-1] == "microchip" {
					floor = append(floor, Item{strings.Replace(info[i+1], "-compatible", "", -1), info[i+2][:len(info[i+2])-1]})
				} else {
					floor = append(floor, Item{info[i+1], info[i+2][:len(info[i+2])-1]})
				}
			}
		}
		floors = append(floors, floor)
	}
	return floors
}

func min_path(f []Floor) int {
	paths := find_moves(State{f, 0, 0})
	for len(paths) > 0 {
		curr_state, paths := paths[len(paths)-1], paths[:len(paths)-1]
		if is_completed(curr_state.floors) {
			return curr_state.distance
		}

		paths = append(paths, find_moves(curr_state)...)
	}
	return -1
}

func find_moves(s State) []State {
	var ret = []State{}
	// TODO
	return ret
}

func is_completed(floors []Floor) bool {
	for i := 0; i < 3; i++ {
		if len(floors[i]) > 0 {
			return false
		}
	}
	return true
}
