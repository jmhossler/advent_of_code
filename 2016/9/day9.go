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

	fmt.Printf("Part 1: decompressed length is %d\n", len(pt1Decompress(data[:len(data)-1])))
	fmt.Printf("Part 2: decompressed length is %d\n", getDecompressedSize(data[:len(data)-1]))
}

func pt1Decompress(input []byte) []byte {
	var decompressed []byte
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			var marker string
			var finishMarker int
			for j := 1; input[j+i] != ')'; j++ {
				marker += string(input[j+i])
				finishMarker = j
			}
			markerInfo := strings.Split(marker, "x")
			length, _ := strconv.Atoi(markerInfo[0])
			repetitions, _ := strconv.Atoi(markerInfo[1])
			repeatedStr := input[i+finishMarker+2 : i+finishMarker+2+length]
			for j := 0; j < repetitions; j++ {
				decompressed = append(decompressed, repeatedStr...)
			}
			i += finishMarker + length + 1
		} else {
			decompressed = append(decompressed, input[i])
		}
	}
	return decompressed
}

func getDecompressedSize(input []byte) int {
	var decompressedLength int
	for i := 0; i < len(input); i++ {
		if input[i] == '(' {
			var marker string
			var finishMarker int
			for j := 1; input[j+i] != ')'; j++ {
				marker += string(input[j+i])
				finishMarker = j
			}
			markerInfo := strings.Split(marker, "x")
			length, _ := strconv.Atoi(markerInfo[0])
			repetitions, _ := strconv.Atoi(markerInfo[1])
			repeatedStr := input[i+finishMarker+2 : i+finishMarker+2+length]
			repeatedLength := getDecompressedSize(repeatedStr)
			for j := 0; j < repetitions; j++ {

				decompressedLength += repeatedLength
			}
			i += length + finishMarker + 1
		} else {
			decompressedLength++
		}
	}
	return decompressedLength
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
