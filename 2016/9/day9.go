package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 9 of Advent of Code 2016")
	data, err := ioutil.ReadFile("input")
	check(err)

	fmt.Printf("Part 1: decompressed length is %d\n", len(pt1_decompress(data[:len(data)-1])))
	fmt.Printf("Part 2: decompressed length is %d\n", get_decompressed_size(data[:len(data)-1]))
}

func pt1_decompress(input []byte) []byte {
	var decompressed []byte
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			var marker string
			var finish_marker int
			for j := 1; input[j+i] != ')'; j++ {
				marker += string(input[j+i])
				finish_marker = j
			}
			marker_info := strings.Split(marker, "x")
			length, _ := strconv.Atoi(marker_info[0])
			repetitions, _ := strconv.Atoi(marker_info[1])
			repeated_str := input[i+finish_marker+2 : i+finish_marker+2+length]
			for j := 0; j < repetitions; j++ {
				decompressed = append(decompressed, repeated_str...)
			}
			i += finish_marker + length + 1
		} else {
			decompressed = append(decompressed, input[i])
		}
	}
	return decompressed
}

func get_decompressed_size(input []byte) int {
	var decompressed_length int
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			var marker string
			var finish_marker int
			for j := 1; input[j+i] != ')'; j++ {
				marker += string(input[j+i])
				finish_marker = j
			}
			marker_info := strings.Split(marker, "x")
			length, _ := strconv.Atoi(marker_info[0])
			repetitions, _ := strconv.Atoi(marker_info[1])
			repeated_str := input[i+finish_marker+2 : i+finish_marker+2+length]
			repeated_length := get_decompressed_size(repeated_str)
			for j := 0; j < repetitions; j++ {

				decompressed_length += repeated_length
			}
			i += length + finish_marker + 1
		} else {
			decompressed_length += 1
		}
	}
	return decompressed_length
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
