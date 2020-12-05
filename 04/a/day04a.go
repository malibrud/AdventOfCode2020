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
	passports := strings.Split(string(dat), "\r\n\r\n")

	validCount := 0
	for _, passport := range passports {
		valid := true
		valid = valid && strings.Contains(passport, "byr:")
		valid = valid && strings.Contains(passport, "iyr:")
		valid = valid && strings.Contains(passport, "eyr:")
		valid = valid && strings.Contains(passport, "hgt:")
		valid = valid && strings.Contains(passport, "hcl:")
		valid = valid && strings.Contains(passport, "ecl:")
		valid = valid && strings.Contains(passport, "pid:")
		if valid {
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
