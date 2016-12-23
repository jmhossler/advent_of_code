package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"strings"
)

type path struct {
	x, y int
	path string
}

var password = flag.String("s", "qtetzkpl", "password")

func main() {
	flag.Parse()
	fmt.Println("Day 17 of Advent of Code 2016")

	path, longest := shortestPath(*password)
	fmt.Printf("Part 1: %s\n", path)
	fmt.Printf("Part 2: %d\n", len(longest))
}

func shortestPath(p string) (string, string) {
	visited := []path{path{0, 0, p}}

	var shortest []string
	for len(visited) > 0 {
		currPath := visited[0]
		visited = visited[1:]

		if currPath.x == 3 && currPath.y == 3 {
			shortest = append(shortest, strings.Replace(currPath.path, p, "", -1))
		} else {

			code := getCode(currPath.path)
			if isOpen(code[0]) && currPath.y > 0 {
				visited = append(visited, path{currPath.x, currPath.y - 1, currPath.path + "U"})
			}
			if isOpen(code[1]) && currPath.y < 3 {
				visited = append(visited, path{currPath.x, currPath.y + 1, currPath.path + "D"})
			}
			if isOpen(code[2]) && currPath.x > 0 {
				visited = append(visited, path{currPath.x - 1, currPath.y, currPath.path + "L"})
			}
			if isOpen(code[3]) && currPath.x < 3 {
				visited = append(visited, path{currPath.x + 1, currPath.y, currPath.path + "R"})
			}
		}
	}
	return shortest[0], shortest[len(shortest)-1]
}

func isOpen(char byte) bool {
	return char == 'b' || char == 'c' || char == 'd' || char == 'e' || char == 'f'
}

func getCode(s string) string {
	str := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	return str[:4]
}
