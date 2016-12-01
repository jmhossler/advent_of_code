package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Coord struct {
	n int
	e int
}

func main() {
	fmt.Println("Day 1 of 2016 Advent of Code challenge")
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
	var first_visited_twice Coord
	is_done := false
	var visited []Coord
	for i := 0; i < len(directions); i++ {
		distance := get_val(strings.TrimSpace(directions[i]))
		direction := get_direction(strings.TrimSpace(directions[i]))

		if curr == north {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					e -= 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = east
			} else {
				for j := 0; j < distance; j++ {
					e += 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = west
			}
		} else if curr == south {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					e += 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = west
			} else {
				for j := 0; j < distance; j++ {
					e -= 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = east
			}
		} else if curr == west {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					n += 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = north
			} else {
				for j := 0; j < distance; j++ {
					n -= 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = south
			}
		} else if curr == east {
			if direction == 'L' {
				for j := 0; j < distance; j++ {
					n -= 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = south
			} else {
				for j := 0; j < distance; j++ {
					n += 1
					if check_exist(Coord{n, e}, visited) && !is_done {
						first_visited_twice = Coord{n, e}
						is_done = true
					}
					visited = append(visited, Coord{n, e})
				}
				curr = north
			}
		}
	}

	fmt.Printf("Distance: %d\n", abs(n)+abs(e))

	fmt.Printf("First visited twice: %d %d\n", first_visited_twice.n, first_visited_twice.e)
	fmt.Printf("Distance: %d\n", abs(first_visited_twice.n)+abs(first_visited_twice.e))
}

func check_exist(c Coord, arr []Coord) bool {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		if temp.n == c.n && temp.e == c.e {
			return true
		}
	}
	return false
}

func get_val(direction string) int {
	val, _ := strconv.Atoi(direction[1:])
	return val
}

func get_direction(direction string) rune {
	return rune(direction[0])
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	} else {
		return val
	}
}
