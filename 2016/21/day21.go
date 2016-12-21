package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var password = flag.String("p", "abcdefgh", "password to scramble")

func main() {
	flag.Parse()
	fmt.Println("Day 21 of Advent of Code 2016")

	input := readInput(os.Stdin)

	scrambled := scrambler([]byte(*password), input)
	fmt.Printf("Part 1: %s\n", scrambled)

	permutations := getPermutations("fbgdceah")
	fmt.Println(permutations)
	for i := 0; i < len(permutations); i++ {
		if scrambler(permutations[i], input) == "fbgdceah" {
			fmt.Printf("Part 2: %s\n", permutations[i])
			break
		}
	}
}

func getPermutations(s string) [][]byte {
	permutations := [][]byte{}
	for i := 0; i < len(s); i++ {
		fmt.Println(strings.Replace(s, string(s[i]), "", -1))
		subPermutations := getPermutations(strings.Replace(s, string(s[i]), "", -1))
		for _, p := range subPermutations {
			substring := make([]byte, len(s))
			substring[0] = s[i]
			for j, b := range p {
				substring[j+1] = b
			}
			permutations = append(permutations, substring)
		}
	}
	return permutations
}

func scrambler(password []byte, commands []string) string {
	for _, cmd := range commands {
		//fmt.Printf("%s:\n", cmd)
		fields := strings.Fields(cmd)
		switch fields[0] {
		case "swap":
			var x, y byte
			if fields[1] == "position" {
				a, _ := strconv.Atoi(fields[2])
				b, _ := strconv.Atoi(fields[5])
				x = password[a]
				y = password[b]
			} else {
				x = fields[2][0]
				y = fields[5][0]
			}
			for i := 0; i < len(password); i++ {
				if password[i] == x {
					password[i] = y
				} else if password[i] == y {
					password[i] = x
				}
			}
		case "rotate":
			dir := fields[1]
			var x int
			if dir == "based" {
				char := fields[6][0]
				dir = "right"
				var index int
				for i := 0; i < len(password); i++ {
					if password[i] == char {
						index = i
					}
				}
				x = index
				if x > 3 {
					x += 1
				}
				x += 1
			} else {
				x, _ = strconv.Atoi(fields[2])
			}
			password = rotatePassword(password, dir, x)
		case "reverse":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[4])
			for i := 0; i < ((y-x)/2)+1; i++ {
				password[x+i], password[y-i] = password[y-i], password[x+i]
			}
		case "move":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[5])
			char := password[x]
			password = append(password[:x], password[x+1:]...)
			password = append(password[:y], append([]byte{char}, password[y:]...)...)
		}
		//fmt.Printf("result:\t%s\n", password)
	}
	return string(password)
}

func rotatePassword(p []byte, dir string, n int) []byte {
	for i := 0; i < n; i++ {
		p = rotateOnce(p, dir)
	}
	return p
}

func rotateOnce(p []byte, dir string) []byte {
	if dir == "left" {
		return append(p[1:], p[0])
	}
	return append(p[len(p)-1:], p[:len(p)-1]...)
}

func readInput(f *os.File) (data []string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return
}
