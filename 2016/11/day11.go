package main

import "fmt"

var floors = [][]string{
	[]string{"PRG", "PM"},
	[]string{"COG", "CUG", "RUG", "PLG"},
	[]string{"COM", "CUM", "RUM", "PLM"},
	[]string{},
}

type Floor struct {
	items []Item
}

type Item struct {
	name        string
	description string
}

var prg = Item{"promethium", "generator"}
var prm = Item{"promethium", "microchip"}
var cog = Item{"cobalt", "generator"}
var cug = Item{"curium", "generator"}
var rug = Item{"ruthenium", "generator"}
var plg = Item{"plutonium", "generator"}
var com = Item{"cobalt", "microchip"}
var cum = Item{"curium", "microchip"}
var rum = Item{"ruthenium", "microchip"}
var plm = Item{"plutonium", "microchip"}

func main() {
	fmt.Println("Day 11 of Advent of Code 2016")
	var floors = [][]Item{[]Item{prg, prm}, []Item{cog, cug, rug, plg}, []Item{com, cum, rum, plm}, []item{}}

	fmt.Println(floors)
	min := make_move(floors, 0)

	fmt.Printf("Part 1: %d\n", min)
}

func make_move(floors [][]Item, floor int) int {
	if is_completed() {
		return 0
	} else {
		min := 100000000
		if floor == 0 {
			moves := find_moves(floors[0], floors[1])
			for _, move := range moves {
				rank := make_move(apply(floors, move, 1), floor+1)
				if rank < min {
					min = rank
				}
			}
		} else {
			moves_down := find_moves(floors[floor], floors[floor-1])
			for _, move := range moves_down {
				rank := make_move(apply(floors, move, -1), floor-1)
				if rank < min {
					min = rank
				}
			}
			moves_up := find_moves(floors[floor], floors[floor+1])
			for _, move := range moves_up {
				rank := make_move(apply(floors, move, 1), floor+1)
				if rank < min {
					min = rank
				}
			}
		}
		return min
	}
}

func apply(floors [][]Item, move []Item, dir int) [][]Item {
	var ret = make([][]Item, 4)
	// TODO
	return ret
}

func find_moves(from, to []Item) [][]Item {
	var ret = [][]Item{}
	// TODO
	return ret
}

func is_completed() bool {
	for i := 0; i < 3; i++ {
		if len(floors[i]) != 1 {
			return false
		}
	}
	return true
}
