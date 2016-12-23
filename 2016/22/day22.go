package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	x    int
	y    int
	size int
	used int
}

func main() {
	fmt.Println("Day 22 of Advent of Code 2016")
	nodes := readInput(os.Stdin)

	viablePairs := countViablePairs(nodes)
	fmt.Printf("Part 1: %d viable pairs\n", viablePairs)
	displayNodes(nodes)
	fmt.Println("Part 2 solved using graph")
}

func countViablePairs(n map[int]map[int]node) int {
	notEmpty := []node{}
	for _, yMap := range n {
		for _, node := range yMap {
			if node.used != 0 {
				notEmpty = append(notEmpty, node)
			}
		}
	}
	var count int
	for _, yMap := range n {
		for _, b := range yMap {
			for _, a := range notEmpty {
				if b != a && a.used <= (b.size-b.used) {
					count++
				}
			}
		}
	}
	return count
}

func displayNodes(n map[int]map[int]node) {
	maxX := getMaxX(n)
	maxY := getMaxY(n)

	fmt.Printf("    ")
	for k := 0; k <= maxX; k++ {
		fmt.Printf("     %-5d", k)
	}
	fmt.Printf("\n")
	for i := 0; i <= maxY; i++ {
		fmt.Printf("%2d: ", i)
		for j := 0; j <= maxX; j++ {
			if n[j][i].used == 0 {
				fmt.Printf("     _    ")
			} else {
				fmt.Printf(" %4d/%-4d", n[j][i].used, n[j][i].size)
			}
			/*
				if isHigh(n, j, i) {
					fmt.Printf(" #")
				} else {
					fmt.Printf(" .")
				}
			*/
		}
		fmt.Printf("\n")
	}
}

func getMaxX(n map[int]map[int]node) int {
	max := 0
	for k := range n {
		if k > max {
			max = k
		}
	}
	return max
}

func getMaxY(n map[int]map[int]node) int {
	max := 0
	for _, sub := range n {
		for k := range sub {
			if k > max {
				max = k
			}
		}
	}
	return max
}

func readInput(f *os.File) map[int]map[int]node {
	scanner := bufio.NewScanner(f)
	nodes := make(map[int]map[int]node)
	scanner.Scan() // skip first line - not important
	scanner.Scan() // skip second line - not important
	for scanner.Scan() {
		info := strings.Fields(scanner.Text())
		name := getName(info[0])
		size, err := strconv.Atoi(info[1][:len(info[1])-1])
		check(err)
		used, err := strconv.Atoi(info[2][:len(info[2])-1])
		check(err)
		getX := regexp.MustCompile(`x\d+`)
		getY := regexp.MustCompile(`y\d+`)
		x, err := strconv.Atoi(getX.FindString(name)[1:])
		check(err)
		y, err := strconv.Atoi(getY.FindString(name)[1:])
		check(err)

		if insideMap, ok := nodes[x]; ok {
			insideMap[y] = node{x, y, size, used}
		} else {
			nodes[x] = make(map[int]node)
			nodes[x][y] = node{x, y, size, used}
		}
	}
	return nodes
}

func getName(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			return s[i+1:]
		}
	}
	return s
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
