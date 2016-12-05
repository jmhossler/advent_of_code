package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Day 5 solution Advent of Code 2016")

	data := []byte("reyedfim")
	hash := md5.Sum(data)

	f := bufio.NewWriter(os.Stdout)

	var password string
	var harder_password [8]byte
	slots := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for index := 0; len(password) < 8 || !slots_filled(slots); index++ {
		str_rep := strconv.Itoa(index)
		new_data := data
		for _, val := range str_rep {
			new_data = append(new_data, byte(val))
		}
		//fmt.Printf("Str_Rep of %s + %d: %s\n", data, index, new_data)
		hash = md5.Sum(new_data)
		str_hash := fmt.Sprintf("%x", hash)
		if check_hash(str_hash) {
			if len(password) < 8 {
				password += string([]byte{str_hash[5]})
			}
			i, err := strconv.Atoi(string([]byte{str_hash[5]}))
			if i < 8 && slots[i] == 0 && err == nil {
				//fmt.Fprintf(f, "%s: ", str_hash)
				harder_password[i] = str_hash[6]
				slots[i] = 1
				display_b_pass(f, harder_password, slots)
				f.Flush()
			}
		}
	}

	fmt.Printf("Part 1: %s\n", password)
	fmt.Printf("Part 2: %s\n", harder_password)
}

func display_b_pass(f *bufio.Writer, p [8]byte, s []int) {
	var str string
	for i := 0; i < 8; i++ {
		if s[i] == 1 {
			str += string([]byte{p[i]})
		} else {
			str += "_"
		}
	}
	if slots_filled(s) {
		fmt.Fprintf(f, "%s\n", str)
	} else {
		fmt.Fprintf(f, "%s\r", str)
	}
}

func slots_filled(s []int) bool {
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

func check_hash(h string) bool {
	return h[:5] == "00000"
}
