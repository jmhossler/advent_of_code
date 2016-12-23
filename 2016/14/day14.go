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
var stretchedHashes map[int]string

func main() {
	flag.Parse()
	fmt.Println("Day 14 of Advent of Code 2016")
	var index int
	var keys []string
	var stretchedKeys []string
	hashes = make(map[int]string)
	stretchedHashes = make(map[int]string)

	for len(keys) < 64 {
		hash := getHash(index)

		if isKey(hashes, getHash, hash, index) && len(keys) < 64 && !contains(keys, hash) {
			keys = append(keys, hash)
		}

		index++
	}
	fmt.Printf("Part 1: %d\n", index-1)

	var sIndex int
	for len(stretchedKeys) < 64 {
		stretchedHash := getStretched(sIndex)

		if isKey(stretchedHashes, getStretched, stretchedHash, sIndex) && len(stretchedKeys) < 64 && !contains(stretchedKeys, stretchedHash) {
			stretchedKeys = append(stretchedKeys, stretchedHash)
		}
		sIndex++
	}

	fmt.Printf("Part 2: %d\n", sIndex-1)
}

func getStretched(index int) string {
	if stretchedHashes[index] == "" {
		hash := getHash(index)
		for i := 0; i < 2016; i++ {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
		}
		stretchedHashes[index] = hash
	}
	return stretchedHashes[index]
}

func getHash(index int) string {
	if hashes[index] == "" {

		hashes[index] = fmt.Sprintf("%x", md5.Sum([]byte(*input+strconv.Itoa(index))))
	}
	return hashes[index]
}

func isKey(src map[int]string, fn func(int) string, hash string, index int) bool {
	if c, err := repeatedThree(hash); err == nil {
		for i := index + 1; i <= index+1000; i++ {
			newHash := fn(i)
			repeated, err := repeatedFive(newHash)
			if err == nil {
				for _, nC := range repeated {
					if c == nC {
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

func repeatedFive(s string) ([]string, error) {
	matches := []string{}
	for i := 0; i < len(s)-4; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] && s[i] == s[i+3] && s[i] == s[i+4] {
			matches = append(matches, string(s[i]))
		}
	}

	if len(matches) == 0 {
		return []string{}, errors.New("Bad")
	}
	return matches, nil
}

func repeatedThree(s string) (string, error) {
	matches := []string{}

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			matches = append(matches, string(s[i]))
		}
	}
	if len(matches) == 0 {
		return "", errors.New("Bad")
	}
	return matches[0], nil

}
