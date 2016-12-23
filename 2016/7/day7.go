package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Day 7 of Advent of Code 2016")

	information := readInput()

	var countTLS, countSSL int
	for i := 0; i < len(information); i++ {
		brackets := extractBrackets(information[i])
		supernet := information[i]
		for _, str := range brackets {
			supernet = strings.Replace(supernet, str, "", -1)
		}
		ABAs := getAbas(supernet)
		var BABs []string
		for _, str := range brackets {
			BABs = append(BABs, getAbas(str[1:len(str)-1])...)
		}
		if validateAbas(ABAs, BABs) {
			countSSL++
		}
		if validBrackets(brackets) {
			if hasAbba(information[i]) {
				countTLS++
			}
		}
	}

	fmt.Printf("Part 1: %d IPs support TLS\n", countTLS)
	fmt.Printf("Part 2: %d IPs support SSL\n", countSSL)
}

func getAbas(str string) []string {
	var abas []string
	for i := 1; i < len(str)-1; i++ {
		if str[i-1] == str[i+1] {
			abas = append(abas, str[i-1:i+2])
		}
	}
	return abas
}

func validateAbas(ABAs, BABs []string) bool {
	for _, aba := range ABAs {
		for _, bab := range BABs {
			if abaMatch(aba, bab) {
				return true
			}
		}
	}
	return false
}

func abaMatch(aba, bab string) bool {
	return aba[1] == bab[0] && aba[0] == bab[1]
}

func validBrackets(b []string) bool {
	for _, str := range b {
		if hasAbba(str) {
			return false
		}
	}
	return true
}

func hasAbba(str string) bool {
	var a, b byte
	a, b = str[0], str[1]
	for i := 2; i < len(str)-1; i++ {
		if b == str[i] && a == str[i+1] && a != b {
			return true
		}
		a, b = str[i-1], str[i]
	}
	return false
}

func extractBrackets(str string) []string {
	re := regexp.MustCompile("\\[[[:alpha:]]+\\]")

	return re.FindAllString(str, -1)
}

func readInput() []string {
	f, err := os.Open("input")
	check(err)

	scanner := bufio.NewScanner(f)

	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
