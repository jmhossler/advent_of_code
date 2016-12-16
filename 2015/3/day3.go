package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	x, y int
}

type CoordArr []Coord

func main() {
	fmt.Println("Advent of Code 2015 day 3")

	data, _ := ioutil.ReadFile("input")
	input := strings.TrimSpace(string(data))

	curr := Coord{0, 0}
	visited := CoordArr{curr}

	for i := 0; i < len(input); i++ {
		if input[i] == '>' {
			curr = Coord{curr.x + 1, curr.y}
		} else if input[i] == '<' {
			curr = Coord{curr.x - 1, curr.y}
		} else if input[i] == '^' {
			curr = Coord{curr.x, curr.y + 1}
		} else if input[i] == 'v' {
			curr = Coord{curr.x, curr.y - 1}
		}
		if !visited.contains(curr) {
			visited = append(visited, curr)
		}
	}

	fmt.Printf("Part one: %d unique houses visited\n", len(visited))

	santa_curr := Coord{0, 0}
	robo_curr := Coord{0, 0}
	visited = CoordArr{santa_curr}
	for i := 0; i < len(input); i += 2 {
		if input[i] == '>' {
			santa_curr = Coord{santa_curr.x + 1, santa_curr.y}
		} else if input[i] == '<' {
			santa_curr = Coord{santa_curr.x - 1, santa_curr.y}
		} else if input[i] == '^' {
			santa_curr = Coord{santa_curr.x, santa_curr.y + 1}
		} else if input[i] == 'v' {
			santa_curr = Coord{santa_curr.x, santa_curr.y - 1}
		}
		if input[i+1] == '>' {
			robo_curr = Coord{robo_curr.x + 1, robo_curr.y}
		} else if input[i+1] == '<' {
			robo_curr = Coord{robo_curr.x - 1, robo_curr.y}
		} else if input[i+1] == '^' {
			robo_curr = Coord{robo_curr.x, robo_curr.y + 1}
		} else if input[i+1] == 'v' {
			robo_curr = Coord{robo_curr.x, robo_curr.y - 1}
		}
		if !visited.contains(santa_curr) {
			visited = append(visited, santa_curr)
		}
		if !visited.contains(robo_curr) {
			visited = append(visited, robo_curr)
		}
	}
	fmt.Printf("Part two: %d unique houses visited\n", len(visited))
}

func (v CoordArr) contains(c Coord) bool {
	for i := 0; i < len(v); i++ {
		if v[i].x == c.x && v[i].y == c.y {
			return true
		}
	}
	return false
}
