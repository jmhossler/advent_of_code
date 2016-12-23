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
		input = stepString(input)
	}
	input = input[:*n]

	chksm := checksum(input)
	fmt.Println(chksm)
}

func stepString(s string) string {
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
	var newString []rune
	for _, c := range s {
		if c == '1' {
			newString = append(newString, '0')
		} else {
			newString = append(newString, '1')
		}
	}
	return string(newString)
}

func reverse(s string) string {
	var newString []byte
	for i := len(s) - 1; i >= 0; i-- {
		newString = append(newString, s[i])
	}
	return string(newString)
}
