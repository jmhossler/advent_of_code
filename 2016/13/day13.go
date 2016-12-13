package main

import (
	"flag"
	"fmt"
)

var input = flag.Int("v", 1352, "value")

func main() {
	flag.Parse()
	fmt.Println("Day 13 of Advent of Code 2016")

	min := bfs()
	fmt.Printf("Part 1: %d\n", min)
	num := bfs2()
	fmt.Printf("Part 2: %d\n", num)
}

func build_graph(a, b int) [][]byte {
	var initial = make([][]byte, a)
	for i := 0; i < a; i++ {
		initial[i] = make([]byte, b)
		for j := 0; j < b; j++ {
			if is_wall(i, j) {
				initial[i][j] = '#'
			} else {
				initial[i][j] = '.'
			}
		}
	}
	return initial
}

func bfs2() int {
	var visited = [][]int{[]int{1, 1}}
	var next = [][]int{[]int{0, 1, 1}}
	curr := []int{0, 1, 1}

	for curr[0] < 51 {
		curr = next[0]
		x, y := curr[1], curr[2]
		next = next[1:]
		for i := -1; i < 2; i++ {
			if !is_wall(x+i, y) && !has_visited(visited, []int{x + i, y}) && x+i >= 0 && y >= 0 {
				visited = append(visited, []int{x + i, y})
				next = append(next, []int{curr[0] + 1, x + i, y})
			}
		}
		for i := -1; i < 2; i++ {
			if !is_wall(x, y+i) && !has_visited(visited, []int{x, y + i}) && x >= 0 && y+i >= 0 {
				visited = append(visited, []int{x, y + i})
				next = append(next, []int{curr[0] + 1, x, y + i})
			}
		}
	}
	return len(visited)
}

func bfs() int {
	var visited = [][]int{[]int{1, 1}}
	var next = [][]int{[]int{0, 1, 1}}
	curr := []int{0, 1, 1}

	for len(next) != 0 {
		curr = next[0]
		if curr[1] == 31 && curr[2] == 39 {
			return curr[0]
		}
		x, y := curr[1], curr[2]
		next = next[1:]
		for i := -1; i < 2; i++ {
			if !is_wall(x+i, y) && !has_visited(visited, []int{x + i, y}) && x+i >= 0 && y >= 0 {
				visited = append(visited, []int{x + i, y})
				next = append(next, []int{curr[0] + 1, x + i, y})
			}
		}
		for i := -1; i < 2; i++ {
			if !is_wall(x, y+i) && !has_visited(visited, []int{x, y + i}) && x >= 0 && y+i >= 0 {
				visited = append(visited, []int{x, y + i})
				next = append(next, []int{curr[0] + 1, x, y + i})
			}
		}
	}
	return -1
}

func has_visited(v [][]int, i []int) bool {
	for _, pair := range v {
		if i[0] == pair[0] && i[1] == pair[1] {
			return true
		}
	}
	return false
}

func is_wall(x, y int) bool {
	val := x*x + 3*x + 2*x*y + y + y*y + *input
	count := 0
	for i := uint(0); i < 64; i++ {
		if ((1 << i) & val) != 0 {
			count += 1
		}
	}
	if count%2 == 0 {
		return false
	}
	return true
}
