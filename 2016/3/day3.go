package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Tri struct {
	x, y, z int
}

func main() {
	fmt.Println("Day 3 of Advent of Code 2016")

	data, _ := ioutil.ReadFile("input")

	input := string(data)
	count := 0

	input = input[:len(input)-1]

	for _, val := range strings.Split(input, "\n") {
		fmt.Println(val)

		x := get_tri(val)

		if x.is_valid() {
			count += 1
		}
	}

	var a, b, c Tri
	new_count := 0
	iter := 0
	for _, val := range strings.Split(input, "\n") {
		vals := strings.Fields(val)
		if iter == 0 {
			a.x, _ = strconv.Atoi(vals[0])
			b.x, _ = strconv.Atoi(vals[1])
			c.x, _ = strconv.Atoi(vals[2])
			iter = 1
		} else if iter == 1 {
			a.y, _ = strconv.Atoi(vals[0])
			b.y, _ = strconv.Atoi(vals[1])
			c.y, _ = strconv.Atoi(vals[2])
			iter = 2
		} else if iter == 2 {
			a.z, _ = strconv.Atoi(vals[0])
			b.z, _ = strconv.Atoi(vals[1])
			c.z, _ = strconv.Atoi(vals[2])
			iter = 0
			if a.is_valid() {
				new_count += 1
			}
			if b.is_valid() {
				new_count += 1
			}
			if c.is_valid() {
				new_count += 1
			}
		}
	}

	fmt.Printf("Part 1 solution: %d\n", count)
	fmt.Printf("Part 2 solution: %d\n", new_count)
}

func (t Tri) is_valid() bool {
	if t.x+t.y <= t.z {
		return false
	}
	if t.x+t.z <= t.y {
		return false
	}
	if t.y+t.z <= t.x {
		return false
	}
	return true
}

func get_tri(line string) Tri {
	split := strings.Fields(line)
	var vals []int
	for _, val := range split {
		new_val, _ := strconv.Atoi(val)
		vals = append(vals, new_val)
	}

	return Tri{vals[0], vals[1], vals[2]}
}
