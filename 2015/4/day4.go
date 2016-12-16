package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"strconv"
	"time"
)

var input = flag.String("s", "iwrupvqb", "puzzle input")

var hashes = make(map[int]string)

func main() {
	flag.Parse()

	start := time.Now()
	var i int
	for i = 0; !isValid(getHash(*input, i)); i++ {
	}
	elapsed := time.Since(start)
	fmt.Printf("Part 1: %d -- Elapsed %s\n", i, elapsed)

	start = time.Now()
	for i = 0; !pt2isValid(getHash(*input, i)); i++ {
	}
	elapsed = time.Since(start)
	fmt.Printf("Part 2: %d -- Elapsed %s\n", i, elapsed)
}

func getHash(s string, i int) string {
	bytes := append([]byte(s), []byte(strconv.Itoa(i))...)
	/*
		hash := hashes[i]
		if hash == "" {
			hash = fmt.Sprintf("%x", md5.Sum(bytes))
			hashes[i] = hash
		}
		return hash
	*/
	//str := s + strconv.Itoa(i)
	return fmt.Sprintf("%x", md5.Sum(bytes))
}

func pt2isValid(h string) bool {
	return h[:6] == "000000"
}

func isValid(h string) bool {
	return h[:5] == "00000"
}
