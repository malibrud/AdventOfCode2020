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
	groups := strings.Split(string(dat), "\r\n\r\n")

	count := 0
	for _, g := range groups {
		count += getGroupCommonCount(g)
	}
	return strconv.Itoa(count)
}

func getGroupCommonCount(g string) int {
	var common [26]bool
	for i := 0; i < 26; i++ {
		common[i] = true
	}

	travelers := strings.Split(g, "\r\n")
	for _, t := range travelers {
		var qs [26]bool
		for _, c := range t {
			if 'a' <= c && c <= 'z' {
				i := int(c - 'a')
				qs[i] = true
			}
		}
		for i := 0; i < 26; i++ {
			common[i] = common[i] && qs[i]
		}
	}
	count := 0
	for i := 0; i < 26; i++ {
		if common[i] {
			count++
		}
	}
	return count
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
