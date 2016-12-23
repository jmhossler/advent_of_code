package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	x, y int
}

type coordArr []coord

func main() {
	fmt.Println("Advent of Code 2015 day 3")

	data, _ := ioutil.ReadFile("input")
	input := strings.TrimSpace(string(data))

	curr := coord{0, 0}
	visited := coordArr{curr}

	for i := 0; i < len(input); i++ {
		if input[i] == '>' {
			curr = coord{curr.x + 1, curr.y}
		} else if input[i] == '<' {
			curr = coord{curr.x - 1, curr.y}
		} else if input[i] == '^' {
			curr = coord{curr.x, curr.y + 1}
		} else if input[i] == 'v' {
			curr = coord{curr.x, curr.y - 1}
		}
		if !visited.contains(curr) {
			visited = append(visited, curr)
		}
	}

	fmt.Printf("Part one: %d unique houses visited\n", len(visited))

	santaCurr := coord{0, 0}
	roboCurr := coord{0, 0}
	visited = coordArr{santaCurr}
	for i := 0; i < len(input); i += 2 {
		if input[i] == '>' {
			santaCurr = coord{santaCurr.x + 1, santaCurr.y}
		} else if input[i] == '<' {
			santaCurr = coord{santaCurr.x - 1, santaCurr.y}
		} else if input[i] == '^' {
			santaCurr = coord{santaCurr.x, santaCurr.y + 1}
		} else if input[i] == 'v' {
			santaCurr = coord{santaCurr.x, santaCurr.y - 1}
		}
		if input[i+1] == '>' {
			roboCurr = coord{roboCurr.x + 1, roboCurr.y}
		} else if input[i+1] == '<' {
			roboCurr = coord{roboCurr.x - 1, roboCurr.y}
		} else if input[i+1] == '^' {
			roboCurr = coord{roboCurr.x, roboCurr.y + 1}
		} else if input[i+1] == 'v' {
			roboCurr = coord{roboCurr.x, roboCurr.y - 1}
		}
		if !visited.contains(santaCurr) {
			visited = append(visited, santaCurr)
		}
		if !visited.contains(roboCurr) {
			visited = append(visited, roboCurr)
		}
	}
	fmt.Printf("Part two: %d unique houses visited\n", len(visited))
}

func (v coordArr) contains(c coord) bool {
	for i := 0; i < len(v); i++ {
		if v[i].x == c.x && v[i].y == c.y {
			return true
		}
	}
	return false
}
