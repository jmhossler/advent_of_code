package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type floor []item

type item struct {
	name        string
	description string
}

type state struct {
	floors            []floor
	distance, floorID int
}

const (
	microchip = "microchip"
	generator = "generator"
)

func main() {
	fmt.Println("Day 11 of Advent of Code 2016")
	var floors = readInput(os.Stdin)

	min := minPath(floors)

	fmt.Printf("Part 1: %d\n", min)
}

func readInput(f *os.File) []floor {
	scanner := bufio.NewScanner(f)
	var floors []floor
	for scanner.Scan() {
		line := scanner.Text()
		var fl floor
		info := strings.Fields(line)
		for i := range info {
			if info[i] == "a" {
				if info[i+2][0] == 'm' {
					fl = append(fl, item{strings.Replace(info[i+1], "-compatible", "", -1), "microchip"})
				} else {
					fl = append(fl, item{info[i+1], "generator"})
				}
			}
		}
		floors = append(floors, fl)
	}
	return floors
}

func minPath(f []floor) int {
	statesVisited := []state{}
	paths := findMoves(state{f, 0, 0}, &statesVisited)
	for len(paths) > 0 {
		currState := paths[0]
		fmt.Println(currState.distance)
		if isCompleted(currState.floors) {
			return currState.distance
		}

		paths = paths[1:]
		paths = append(paths, findMoves(currState, &statesVisited)...)
	}
	return -1
}

func findMoves(s state, v *[]state) []state {
	var ret = []state{}
	floors := s.floors
	for i := 0; i < len(floors[s.floorID]); i++ {
		var newFloorUp floor
		var newFloorDown floor
		if s.floorID+1 < len(floors) {
			newFloorUp = append(floors[s.floorID+1], floors[s.floorID][i])
			if isValidFloor(newFloorUp) {
				newFloors := floors
				newFloors[s.floorID+1] = newFloorUp
				newState := state{newFloors, s.distance + 1, s.floorID + 1}
				if !hasState(*v, newState) {
					ret = append(ret, newState)
					*v = append(*v, newState)
				}
			}
		}
		if s.floorID-1 >= 0 {
			newFloorDown = append(floors[s.floorID-1], floors[s.floorID][i])
			if isValidFloor(newFloorDown) {
				newFloors := floors
				newFloors[s.floorID-1] = newFloorDown
				newState := state{newFloors, s.distance + 1, s.floorID - 1}
				if !hasState(*v, newState) {
					ret = append(ret, newState)
					*v = append(*v, newState)
				}
			}
		}
		for j := i + 1; j < len(floors[s.floorID]); j++ {
			if s.floorID+1 < len(floors) {
				newFloorUp = append(newFloorUp, floors[s.floorID][j])
				if isValidFloor(newFloorUp) {
					newFloors := floors
					newFloors[s.floorID+1] = newFloorUp
					newState := state{newFloors, s.distance + 1, s.floorID + 1}
					if !hasState(*v, newState) {
						ret = append(ret, newState)
						*v = append(*v, newState)
					}
				}
			}
			if s.floorID-1 >= 0 {
				newFloorDown = append(newFloorDown, floors[s.floorID][j])
				if isValidFloor(newFloorDown) {
					newFloors := floors
					newFloors[s.floorID-1] = newFloorDown
					newState := state{newFloors, s.distance + 1, s.floorID - 1}
					if !hasState(*v, newState) {
						ret = append(ret, newState)
						*v = append(*v, newState)
					}
				}
			}
		}
	}
	return ret
}

func hasState(states []state, s state) bool {
	for _, st := range states {
		if st.floorID == s.floorID {
			if len(st.floors) == len(s.floors) {
				if equalFloors(st.floors, s.floors) {
					return true
				}
			}
		}
	}
	return false
}

func equalFloors(a []floor, b []floor) bool {
	for i := 0; i < len(a); i++ {
		if !equalFloor(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalFloor(a floor, b floor) bool {
	if len(a) == len(b) {
		for _, it := range a {
			if !hasItem(b, it) {
				return false
			}
		}
		return true
	}
	return false
}

func hasItem(f floor, i item) bool {
	for _, it := range f {
		if it.name == i.name && it.description == i.description {
			return true
		}
	}
	return false
}

func isValidFloor(fl floor) bool {
	fl = removeMatches(fl)
	return !(hasMicrochip(fl) && hasGenerator(fl))
}

func removeMatches(f floor) floor {
	var newFloor floor
	for _, it := range f {
		if !hasMatch(f, it) {
			newFloor = append(newFloor, it)
		}
	}
	return newFloor
}

func hasMatch(f floor, i item) bool {
	for _, it := range f {
		if isMatch(i, it) {
			return true
		}
	}
	return false
}

func isMatch(a item, b item) bool {
	if a.name == b.name {
		if a.description == generator {
			return b.description == generator
		}
		return b.description == microchip

	}
	return false
}

func hasMicrochip(f floor) bool {
	for _, it := range f {
		if it.description == microchip {
			return true
		}
	}
	return false
}

func hasGenerator(f floor) bool {
	for _, it := range f {
		if it.description == generator {
			return true
		}
	}
	return false
}

func isCompleted(floors []floor) bool {
	for i := 0; i < 3; i++ {
		if len(floors[i]) > 0 {
			return false
		}
	}
	return true
}
