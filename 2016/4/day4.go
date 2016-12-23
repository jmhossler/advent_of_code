package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type room struct {
	encryptedName string
	id            int
	checksum      string
}

func main() {
	fmt.Println("Day 4 of Advent of Code 2016")

	fp, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(fp)

	var rooms []room

	for scanner.Scan() {
		line := scanner.Text()
		newRoom := room{getName(line), getID(line), getChecksum(line)}
		//fmt.Println(get_checksum(line))
		//fmt.Println(get_id(line))
		//fmt.Println(get_name(line))
		rooms = append(rooms, newRoom)
	}

	sum := 0
	var validRooms []room
	for _, room := range rooms {
		if room.isValid() {
			validRooms = append(validRooms, room)
			sum += room.id
		}
	}

	fmt.Printf("Part 1: sum of valid rooms %d\n", sum)

	fmt.Println("Number of valid rooms: ", len(validRooms))
	for _, room := range validRooms {
		unencryptedName := rotateString(room.encryptedName, room.id)
		fmt.Println(unencryptedName, room.id)
	}
}

func (r room) isValid() bool {
	generatedChecksum := createChecksum(r.encryptedName)
	//fmt.Printf("Gen: %s\tGiven: %s, ans %v\n", generated_checksum, r.checksum, generated_checksum == r.checksum)
	return generatedChecksum == r.checksum
}

func rotateString(str string, n int) string {
	shift, offset := rune(n%26), rune(26)

	runes := []rune(str)
	for i, val := range runes {
		if val >= 'a' && val <= 'z'-shift {
			val = val + shift
		} else if val > 'z'-shift && val <= 'z' {
			val = val + shift - offset
		} else if val == '-' {
			val = ' '
		}
		runes[i] = val
	}

	return string(runes)
}

func createChecksum(name string) string {
	str := strings.Replace(name, "-", "", -1)

	var m = make(map[rune]int)

	for _, val := range str {
		m[val]++
	}

	checksum := ""
	for i := 0; i < 5; i++ {
		k := getMax(m)
		checksum += string(k)
		delete(m, k)
	}
	//fmt.Println("Checksum: ", checksum)

	return checksum
}

func getMax(m map[rune]int) rune {
	var max rune
	for k := range m {
		if m[k] > m[max] {
			max = k
		} else if m[k] == m[max] {
			if k < max {
				max = k
			}
		}
	}
	return max
}

func getChecksum(line string) string {
	re := regexp.MustCompile("\\[[[:alpha:]]{5}\\]")

	str := re.FindString(line)
	return str[1 : len(str)-1]
}

func getID(line string) int {
	re := regexp.MustCompile("-\\d+")

	str := re.FindString(line)
	id, _ := strconv.Atoi(str[1:])
	return id
}

func getName(line string) string {
	re := regexp.MustCompile("([[:alpha:]]+\\-)+")

	str := re.FindString(line)
	return str[:len(str)-1]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
