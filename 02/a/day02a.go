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
		letter := parts[1]

		// Separate lower count from upper count
		parts = strings.Split(counts, "-")
		lc, _ := strconv.Atoi(parts[0])
		uc, _ := strconv.Atoi(parts[1])

		// determine validity
		n := strings.Count(pass, letter)
		if lc <= n && n <= uc {
			validCount++
		}
	}
	return strconv.Itoa(validCount)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
