package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 5 of Advent of Code 2016")

	data := []byte("reyedfim")

	f := bufio.NewWriter(os.Stdout)

	var password []byte
	var harderPassword [8]byte
	slots := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for index := 0; len(password) < 8 || !slotsFilled(slots); index++ {
		strRep := strconv.Itoa(index)
		newData := append(data, []byte(strRep)...)
		hash := md5.Sum(newData)
		strHash := fmt.Sprintf("%x", hash)
		if checkHash(strHash) {
			if len(password) < 8 {
				password = append(password, byte(strHash[5]))
			}
			i, err := strconv.Atoi(string([]byte{strHash[5]}))
			if i < 8 && slots[i] == 0 && err == nil {
				//fmt.Fprintf(f, "%s: ", str_hash)
				harderPassword[i] = strHash[6]
				slots[i] = 1
				displayBPass(f, harderPassword, slots)
				f.Flush()
			}
		}
	}

	fmt.Printf("Part 1: %s\n", password)
	fmt.Printf("Part 2: %s\n", harderPassword)
}

func displayBPass(f *bufio.Writer, p [8]byte, s []int) {
	var str string
	for i := 0; i < 8; i++ {
		if s[i] == 1 {
			str += string([]byte{p[i]})
		} else {
			str += "_"
		}
	}
	if slotsFilled(s) {
		fmt.Fprintf(f, "%s\n", str)
	} else {
		fmt.Fprintf(f, "%s\r", str)
	}
}

func slotsFilled(s []int) bool {
	for _, val := range s {
		if val == 0 {
			return false
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkHash(h string) bool {
	return h[:5] == "00000"
}
