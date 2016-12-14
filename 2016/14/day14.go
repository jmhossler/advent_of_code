package main

import (
	"crypto/md5"
	"errors"
	"flag"
	"fmt"
	"strconv"
)

var input = flag.String("s", "cuanljph", "string input")

var hashes map[int]string
var stretched_hashes map[int]string

func main() {
	flag.Parse()
	fmt.Println("Day 14 of Advent of Code 2016")
	var index int
	var keys []string
	var stretched_keys []string
	hashes = make(map[int]string)
	stretched_hashes = make(map[int]string)

	for len(keys) < 64 {
		hash := get_hash(index)

		if is_key(hashes, get_hash, hash, index) && len(keys) < 64 && !contains(keys, hash) {
			keys = append(keys, hash)
		}

		index += 1
	}
	fmt.Printf("Part 1: %d\n", index-1)

	var s_index int
	for len(stretched_keys) < 64 {
		stretched_hash := get_stretched(s_index)

		if is_key(stretched_hashes, get_stretched, stretched_hash, s_index) && len(stretched_keys) < 64 && !contains(stretched_keys, stretched_hash) {
			stretched_keys = append(stretched_keys, stretched_hash)
		}
		s_index += 1
	}

	fmt.Printf("Part 2: %d\n", s_index-1)
}

func get_stretched(index int) string {
	if stretched_hashes[index] == "" {
		hash := get_hash(index)
		for i := 0; i < 2016; i++ {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
		}
		stretched_hashes[index] = hash
	}
	return stretched_hashes[index]
}

func get_hash(index int) string {
	if hashes[index] == "" {

		hashes[index] = fmt.Sprintf("%x", md5.Sum([]byte(*input+strconv.Itoa(index))))
	}
	return hashes[index]
}

func is_key(src map[int]string, fn func(int) string, hash string, index int) bool {
	if c, err := repeated_three(hash); err == nil {
		for i := index + 1; i <= index+1000; i++ {
			new_hash := fn(i)
			repeated, err := repeated_five(new_hash)
			if err == nil {
				for _, n_c := range repeated {
					if c == n_c {
						return true
					}
				}
			}
		}
	}
	return false
}

func contains(arr []string, s string) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func repeated_five(s string) ([]string, error) {
	matches := []string{}
	for i := 0; i < len(s)-4; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] && s[i] == s[i+3] && s[i] == s[i+4] {
			matches = append(matches, string(s[i]))
		}
	}

	if len(matches) == 0 {
		return []string{}, errors.New("Bad")
	} else {
		return matches, nil
	}
}

func repeated_three(s string) (string, error) {
	matches := []string{}

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			matches = append(matches, string(s[i]))
		}
	}
	if len(matches) == 0 {
		return "", errors.New("Bad")
	} else {
		return matches[0], nil
	}
}
