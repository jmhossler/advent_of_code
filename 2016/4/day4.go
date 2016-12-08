package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Room struct {
	encrypted_name string
	id             int
	checksum       string
}

func main() {
	fmt.Println("Day 4 of Advent of Code 2016")

	fp, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(fp)

	var rooms []Room

	for scanner.Scan() {
		line := scanner.Text()
		new_room := Room{get_name(line), get_id(line), get_checksum(line)}
		//fmt.Println(get_checksum(line))
		//fmt.Println(get_id(line))
		//fmt.Println(get_name(line))
		rooms = append(rooms, new_room)
	}

	sum := 0
	var valid_rooms []Room
	for _, room := range rooms {
		if room.is_valid() {
			valid_rooms = append(valid_rooms, room)
			sum += room.id
		}
	}

	fmt.Printf("Part 1: sum of valid rooms %d\n", sum)

	fmt.Println("Number of valid rooms: ", len(valid_rooms))
	for _, room := range valid_rooms {
		unencrypted_name := rotate_string(room.encrypted_name, room.id)
		fmt.Println(unencrypted_name, room.id)
	}
}

func (r Room) is_valid() bool {
	generated_checksum := create_checksum(r.encrypted_name)
	//fmt.Printf("Gen: %s\tGiven: %s, ans %v\n", generated_checksum, r.checksum, generated_checksum == r.checksum)
	return generated_checksum == r.checksum
}

func rotate_string(str string, n int) string {
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

func create_checksum(name string) string {
	str := strings.Replace(name, "-", "", -1)

	var m map[rune]int = make(map[rune]int)

	for _, val := range str {
		m[val] += 1
	}

	checksum := ""
	for i := 0; i < 5; i++ {
		k := get_max(m)
		checksum += string(k)
		delete(m, k)
	}
	//fmt.Println("Checksum: ", checksum)

	return checksum
}

func get_max(m map[rune]int) rune {
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

func get_checksum(line string) string {
	re := regexp.MustCompile("\\[[[:alpha:]]{5}\\]")

	str := re.FindString(line)
	return str[1 : len(str)-1]
}

func get_id(line string) int {
	re := regexp.MustCompile("-\\d+")

	str := re.FindString(line)
	id, _ := strconv.Atoi(str[1:])
	return id
}

func get_name(line string) string {
	re := regexp.MustCompile("([[:alpha:]]+\\-)+")

	str := re.FindString(line)
	return str[:len(str)-1]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
