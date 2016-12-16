package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n = flag.Int("n", 272, "length")

func main() {
	flag.Parse()
	fmt.Println("Day 16 of Advent of Code 2016")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	for len(input) < *n {
		input = step_string(input)
	}
	input = input[:*n]

	chksm := checksum(input)
	fmt.Println(chksm)
}

func step_string(s string) string {
	return s + "0" + invert(reverse(s))
}

func checksum(s string) string {
	if len(s)%2 == 1 {
		return s
	}
	var chksm []rune
	for i := 0; i < len(s); i += 2 {
		if s[i] == s[i+1] {
			chksm = append(chksm, '1')
		} else {
			chksm = append(chksm, '0')
		}
	}
	return checksum(string(chksm))
}

func invert(s string) string {
	var new_string []rune
	for _, c := range s {
		if c == '1' {
			new_string = append(new_string, '0')
		} else {
			new_string = append(new_string, '1')
		}
	}
	return string(new_string)
}

func reverse(s string) string {
	var new_string []byte
	for i := len(s) - 1; i >= 0; i-- {
		new_string = append(new_string, s[i])
	}
	return string(new_string)
}
