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

	information := read_input()

	var count_tls, count_ssl int
	for i := 0; i < len(information); i++ {
		brackets := extract_brackets(information[i])
		supernet := information[i]
		for _, str := range brackets {
			supernet = strings.Replace(supernet, str, "", -1)
		}
		ABAs := get_abas(supernet)
		var BABs []string
		for _, str := range brackets {
			BABs = append(BABs, get_abas(str[1:len(str)-1])...)
		}
		if validate_abas(ABAs, BABs) {
			count_ssl += 1
		}
		if valid_brackets(brackets) {
			if has_abba(information[i]) {
				count_tls += 1
			}
		}
	}

	fmt.Printf("Part 1: %d IPs support TLS\n", count_tls)
	fmt.Printf("Part 2: %d IPs support SSL\n", count_ssl)
}

func get_abas(str string) []string {
	var abas []string
	for i := 1; i < len(str)-1; i++ {
		if str[i-1] == str[i+1] {
			abas = append(abas, str[i-1:i+2])
		}
	}
	return abas
}

func validate_abas(ABAs, BABs []string) bool {
	for _, aba := range ABAs {
		for _, bab := range BABs {
			if aba_match(aba, bab) {
				return true
			}
		}
	}
	return false
}

func aba_match(aba, bab string) bool {
	return aba[1] == bab[0] && aba[0] == bab[1]
}

func valid_brackets(b []string) bool {
	for _, str := range b {
		if has_abba(str) {
			return false
		}
	}
	return true
}

func has_abba(str string) bool {
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

func extract_brackets(str string) []string {
	re := regexp.MustCompile("\\[[[:alpha:]]+\\]")

	return re.FindAllString(str, -1)
}

func read_input() []string {
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
