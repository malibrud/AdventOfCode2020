package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")

	validCount := 0
	for _, line := range lines {
		// Separate policy from password
		parts := strings.Split(line, ": ")
		policy := parts[0]
		pass := parts[1]

		// Separate policy counts from letter
		parts = strings.Split(policy, " ")
		counts := parts[0]
		c := parts[1][0]

		// Separate lower count from upper count
		parts = strings.Split(counts, "-")
		i1, _ := strconv.Atoi(parts[0])
		i2, _ := strconv.Atoi(parts[1])

		// determine validity
		c1 := pass[i1-1]
		c2 := pass[i2-1]
		if xor(c1 == c, c2 == c) {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func xor(b1 bool, b2 bool) bool {
	return b1 != b2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
