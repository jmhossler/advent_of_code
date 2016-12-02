package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

var keypad [][]int

var one rune = '1'
var two rune = '2'
var three rune = '3'
var four rune = '4'
var five rune = '5'
var six = '6'
var seven = '7'
var eight = '8'
var nine = '9'
var A = 'A'
var B = 'B'
var C rune = 'C'
var D rune = 'D'

var new_keypad = [][]*rune{[]*rune{nil, nil, &one, nil, nil}, []*rune{nil, &two, &three, &four, nil}, []*rune{&five, &six, &seven, &eight, &nine}, []*rune{nil, &A, &B, &C, nil}, []*rune{nil, nil, &D, nil, nil}}

func main() {
	fmt.Println("vim-go")

	keypad = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}

	data, _ := ioutil.ReadFile("input")

	input := strings.Split(string(data), "\n")

	var code string
	var new_code string
	curr_pos := Coord{1, 1}
	pt_2_pos := Coord{0, 2}
	for i := 0; i < len(input)-1; i++ {
		final_coord := get_val(input[i], curr_pos)
		new_final_coord := mod_get_val(input[i], pt_2_pos)

		code += coord_to_val(final_coord)
		new_code += mod_coord_to_val(new_final_coord)
		curr_pos = final_coord
		pt_2_pos = new_final_coord
	}

	fmt.Printf("Part 1: Bathroom code %s\n", code)
	fmt.Printf("Part 2: Bathroom code %s\n", new_code)
}

func get_val(input string, start_pos Coord) Coord {
	for i := 0; i < len(input); i++ {
		if input[i] == 'U' {
			start_pos.y = start_pos.y - 1
			if start_pos.y < 0 {
				start_pos.y = 0
			}
		} else if input[i] == 'D' {
			start_pos.y = start_pos.y + 1
			if start_pos.y > 2 {
				start_pos.y = 2
			}
		} else if input[i] == 'L' {
			start_pos.x = start_pos.x - 1
			if start_pos.x < 0 {
				start_pos.x = 0
			}
		} else if input[i] == 'R' {
			start_pos.x = start_pos.x + 1
			if start_pos.x > 2 {
				start_pos.x = 2
			}
		}
	}

	return start_pos
}

func mod_get_val(input string, start_pos Coord) Coord {
	for i := 0; i < len(input); i++ {
		d_coord := Coord{0, 0}
		if input[i] == 'U' {
			d_coord.y = -1
		} else if input[i] == 'D' {
			d_coord.y = 1
		} else if input[i] == 'L' {
			d_coord.x = -1
		} else if input[i] == 'R' {
			d_coord.x = 1
		}
		new_coord := Coord{start_pos.x + d_coord.x, start_pos.y + d_coord.y}
		if new_coord.x < 5 && new_coord.x >= 0 && new_coord.y < 5 && new_coord.y >= 0 {
			if new_keypad[new_coord.y][new_coord.x] != nil {
				start_pos = new_coord
			}
		}
	}
	return start_pos
}

func coord_to_val(c Coord) string {
	return strconv.Itoa(keypad[c.y][c.x])
}

func mod_coord_to_val(c Coord) string {
	return string(*new_keypad[c.y][c.x])
}
