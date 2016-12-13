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
				if info[i+2][0] == 'm' {
					floor = append(floor, Item{strings.Replace(info[i+1], "-compatible", "", -1), "microchip"})
				} else {
					floor = append(floor, Item{info[i+1], "generator"})
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
		curr_state := paths[0]
		fmt.Println(curr_state.distance)
		if is_completed(curr_state.floors) {
			return curr_state.distance
		}

		paths = paths[1:]
		paths = append(paths, find_moves(curr_state)...)
	}
	return -1
}

func find_moves(s State) []State {
	var ret = []State{}
	floors := s.floors
	for i := 0; i < len(floors[s.floor]); i++ {
		var new_floor_up Floor
		var new_floor_down Floor
		if s.floor+1 < len(floors) {
			new_floor_up = append(floors[s.floor+1], floors[s.floor][i])
			if is_valid_floor(new_floor_up) {
				new_floors := floors
				new_floors[s.floor+1] = new_floor_up
				ret = append(ret, State{new_floors, s.distance + 1, s.floor + 1})
			}
		}
		if s.floor-1 >= 0 {
			new_floor_down = append(floors[s.floor-1], floors[s.floor][i])
			if is_valid_floor(new_floor_down) {
				new_floors := floors
				new_floors[s.floor-1] = new_floor_down
				ret = append(ret, State{new_floors, s.distance + 1, s.floor - 1})
			}
		}
		for j := i + 1; j < len(floors[s.floor]); j++ {
			if s.floor+1 < len(floors) {
				new_floor_up = append(new_floor_up, floors[s.floor][j])
				if is_valid_floor(new_floor_up) {
					new_floors := floors
					new_floors[s.floor+1] = new_floor_up
					ret = append(ret, State{new_floors, s.distance + 1, s.floor + 1})
				}
			}
			if s.floor-1 >= 0 {
				new_floor_down = append(new_floor_down, floors[s.floor][j])
				if is_valid_floor(new_floor_down) {
					new_floors := floors
					new_floors[s.floor-1] = new_floor_down
					ret = append(ret, State{new_floors, s.distance + 1, s.floor - 1})
				}
			}
		}
	}
	return ret
}

func is_valid_floor(floor Floor) bool {
	floor = remove_matches(floor)
	return !(has_microchip(floor) && has_generator(floor))
}

func remove_matches(f Floor) Floor {
	var new_floor Floor
	for _, item := range f {
		if !has_match(f, item) {
			new_floor = append(new_floor, item)
		}
	}
	return new_floor
}

func has_match(f Floor, i Item) bool {
	for _, item := range f {
		if is_match(i, item) {
			return true
		}
	}
	return false
}

func is_match(a Item, b Item) bool {
	if a.name == b.name {
		if a.description == "microchip" {
			return b.description == "generator"
		} else {
			return b.description == "microchip"
		}
	}
	return false
}

func has_microchip(f Floor) bool {
	for _, item := range f {
		if item.description == "microchip" {
			return true
		}
	}
	return false
}

func has_generator(f Floor) bool {
	for _, item := range f {
		if item.description == "generator" {
			return true
		}
	}
	return false
}

func is_completed(floors []Floor) bool {
	for i := 0; i < 3; i++ {
		if len(floors[i]) > 0 {
			return false
		}
	}
	return true
}
