package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2 of 2016 Advent of Code")

	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)

	dimensions := strings.Split(str, "\n")

	total_surface_area := 0
	ribbon_length := 0
	for i := 0; i < len(dimensions)-1; i++ {
		d := strings.Split(dimensions[i], "x")
		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])
		total_surface_area += calculate_surface_area(l, w, h)
		ribbon_length += calculate_ribbon_length(l, w, h)
	}

	fmt.Printf("Part 1: Total square feet of wrapping paper is %d\n", total_surface_area)
	fmt.Printf("Part 2: Total length of ribbon is %d\n", ribbon_length)
}

func calculate_surface_area(l int, w int, h int) int {
	lw := l * w
	wh := w * h
	hl := h * l
	return 2*(lw+wh+hl) + min(lw, min(wh, hl))
}

func calculate_ribbon_length(l int, w int, h int) int {
	shortest_perimeter := min(2*l+2*w, min(2*w+2*h, 2*l+2*h))
	return l*w*h + shortest_perimeter
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
