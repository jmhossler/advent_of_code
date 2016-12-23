package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type coord struct {
	n int
	e int
}

func main() {
	fmt.Println("Day 1 of Advent of Code 2016")
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)

	directions := strings.Split(str, ",")
	n, e := 0, 0

	const (
		north = "North"
		south = "South"
		east  = "East"
		west  = "West"
	)

	curr := north
	var firstVisitedTwice coord
	isDone := false
	var visited []coord
	for i := 0; i < len(directions); i++ {
		distance := getVal(strings.TrimSpace(directions[i]))
		direction := getDirection(strings.TrimSpace(directions[i]))

		if curr == north {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					e--
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = east
			} else {
				for j := 0; j < distance; j++ {
					e++
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = west
			}
		} else if curr == south {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					e++
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = west
			} else {
				for j := 0; j < distance; j++ {
					e--
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = east
			}
		} else if curr == west {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					n++
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = north
			} else {
				for j := 0; j < distance; j++ {
					n--
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = south
			}
		} else if curr == east {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					n--
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = south
			} else {
				for j := 0; j < distance; j++ {
					n++
					if checkExist(coord{n, e}, visited) && !isDone {
						firstVisitedTwice = coord{n, e}
						isDone = true
					}
					visited = append(visited, coord{n, e})
				}
				curr = north
			}
		}
	}

	fmt.Println("pt1 Correct answer: 271")
	fmt.Printf("Distance: %d\n\n", abs(n)+abs(e))

	fmt.Println("pt2 Correct answer: 153")
	fmt.Printf("First visited twice: %d %d\n", firstVisitedTwice.n, firstVisitedTwice.e)
	fmt.Printf("Distance: %d\n", abs(firstVisitedTwice.n)+abs(firstVisitedTwice.e))
}

func checkExist(c coord, arr []coord) bool {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if temp.n == c.n && temp.e == c.e {
			return true
		}
	}
	return false
}

func getVal(direction string) int {
	val, _ := strconv.Atoi(direction[1:])
	return val
}

func getDirection(direction string) rune {
	return rune(direction[0])
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}
