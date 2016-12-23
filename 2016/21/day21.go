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
var scrambledPass = flag.String("s", "fbgdceah", "scrambled password")

var reverseMap = map[int]int{
	1: 0,
	3: 1,
	5: 2,
	7: 3,
	2: 4,
	4: 5,
	6: 6,
	0: 7,
}

const (
	left  = "left"
	right = "right"
	based = "based"
)

func main() {
	flag.Parse()
	fmt.Println("Day 21 of Advent of Code 2016")

	input := readInput(os.Stdin)

	scrambled := scrambler([]byte(*password), input)
	fmt.Printf("Part 1: %s\n", scrambled)

	unscrambled := unScramble([]byte(*scrambledPass), input)
	fmt.Printf("Part 2: %s\n", unscrambled)
}

func unScramble(scrambled []byte, commands []string) string {
	for i := len(commands) - 1; i >= 0; i-- {
		fields := strings.Fields(commands[i])
		switch fields[0] {
		case "swap":
			scrambled = swapPass(scrambled, fields)
		case "rotate":
			dir := fields[1]
			var x int
			if dir == based {
				char := fields[6][0]
				dir = left
				index := getIndex(scrambled, char)
				for getIndex(scrambled, char) != reverseMap[index] {
					scrambled = rotatePassword(scrambled, dir, 1)
				}
			} else {
				if dir == left {
					dir = right
				} else {
					dir = left
				}
				x, _ = strconv.Atoi(fields[2])
				scrambled = rotatePassword(scrambled, dir, x)
			}
		case "reverse":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[4])
			scrambled = reversePass(scrambled, x, y)
		case "move":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[5])
			scrambled = movePass(scrambled, y, x)
		}
	}
	return string(scrambled)
}

func diff(a, b int) int {
	result := a - b
	if result < 0 {
		return result * -1
	}
	return result
}

func swapPass(p []byte, fields []string) []byte {
	var x, y byte
	if fields[1] == "position" {
		a, _ := strconv.Atoi(fields[2])
		b, _ := strconv.Atoi(fields[5])
		x = p[a]
		y = p[b]
	} else {
		x = fields[2][0]
		y = fields[5][0]
	}
	for i := 0; i < len(p); i++ {
		if p[i] == x {
			p[i] = y
		} else if p[i] == y {
			p[i] = x
		}
	}
	return p
}

func getIndex(p []byte, c byte) int {
	for i := 0; i < len(p); i++ {
		if p[i] == c {
			return i
		}
	}
	return 0
}

func scrambler(password []byte, commands []string) string {
	for _, cmd := range commands {
		//fmt.Printf("%s:\n", cmd)
		fields := strings.Fields(cmd)
		switch fields[0] {
		case "swap":
			password = swapPass(password, fields)
		case "rotate":
			dir := fields[1]
			var x int
			if dir == based {
				char := fields[6][0]
				dir = right
				x = getIndex(password, char)
				if x > 3 {
					x++
				}
				x++
			} else {
				x, _ = strconv.Atoi(fields[2])
			}
			password = rotatePassword(password, dir, x)
		case "reverse":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[4])
			password = reversePass(password, x, y)
		case "move":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[5])
			password = movePass(password, x, y)
		}
		//fmt.Printf("result:\t%s\n", password)
	}
	return string(password)
}

func reversePass(p []byte, x, y int) []byte {
	for i := 0; i < ((y-x)/2)+1; i++ {
		p[x+i], p[y-i] = p[y-i], p[x+i]
	}
	return p
}

func movePass(p []byte, x, y int) []byte {
	char := p[x]
	p = append(p[:x], p[x+1:]...)
	p = append(p[:y], append([]byte{char}, p[y:]...)...)
	return p
}

func rotatePassword(p []byte, dir string, n int) []byte {
	for i := 0; i < n; i++ {
		p = rotateOnce(p, dir)
	}
	return p
}

func rotateOnce(p []byte, dir string) []byte {
	if dir == left {
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
