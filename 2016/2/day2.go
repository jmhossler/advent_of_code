package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

var keypad [][]int

var one = '1'
var two = '2'
var three = '3'
var four = '4'
var five = '5'
var six = '6'
var seven = '7'
var eight = '8'
var nine = '9'
var a = 'A'
var b = 'B'
var c = 'C'
var d = 'D'

var newKeypad = [][]*rune{[]*rune{nil, nil, &one, nil, nil}, []*rune{nil, &two, &three, &four, nil}, []*rune{&five, &six, &seven, &eight, &nine}, []*rune{nil, &a, &b, &c, nil}, []*rune{nil, nil, &d, nil, nil}}

func main() {
	fmt.Println("Day 2 of Advent of Code 2016")

	keypad = [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}

	data, _ := ioutil.ReadFile("input")

	input := strings.Split(string(data), "\n")

	var code string
	var newCode string
	currPos := coord{1, 1}
	pt2Pos := coord{0, 2}
	for i := 0; i < len(input)-1; i++ {
		finalCoord := getVal(input[i], currPos)
		newFinalCoord := modGetVal(input[i], pt2Pos)

		code += coordToVal(finalCoord)
		newCode += modCoordToVal(newFinalCoord)
		currPos = finalCoord
		pt2Pos = newFinalCoord
	}

	fmt.Printf("Part 1: Bathroom code %s\n", code)
	fmt.Printf("Part 2: Bathroom code %s\n", newCode)
}

func getVal(input string, startPos coord) coord {
	for i := 0; i < len(input); i++ {
		if input[i] == 'U' {
			startPos.y = startPos.y - 1
			if startPos.y < 0 {
				startPos.y = 0
			}
		} else if input[i] == 'D' {
			startPos.y = startPos.y + 1
			if startPos.y > 2 {
				startPos.y = 2
			}
		} else if input[i] == 'L' {
			startPos.x = startPos.x - 1
			if startPos.x < 0 {
				startPos.x = 0
			}
		} else if input[i] == 'R' {
			startPos.x = startPos.x + 1
			if startPos.x > 2 {
				startPos.x = 2
			}
		}
	}

	return startPos
}

func modGetVal(input string, startPos coord) coord {
	for i := 0; i < len(input); i++ {
		dCoord := coord{0, 0}
		if input[i] == 'U' {
			dCoord.y = -1
		} else if input[i] == 'D' {
			dCoord.y = 1
		} else if input[i] == 'L' {
			dCoord.x = -1
		} else if input[i] == 'R' {
			dCoord.x = 1
		}
		newCoord := coord{startPos.x + dCoord.x, startPos.y + dCoord.y}
		if newCoord.x < 5 && newCoord.x >= 0 && newCoord.y < 5 && newCoord.y >= 0 {
			if newKeypad[newCoord.y][newCoord.x] != nil {
				startPos = newCoord
			}
		}
	}
	return startPos
}

func coordToVal(c coord) string {
	return strconv.Itoa(keypad[c.y][c.x])
}

func modCoordToVal(c coord) string {
	return string(*newKeypad[c.y][c.x])
}
