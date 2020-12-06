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
		count += getGroupCount(g)
	}
	return strconv.Itoa(count)
}

func getGroupCount(g string) int {
	count := 0
	var qs [26]bool
	for _, c := range g {
		if 'a' <= c && c <= 'z' {
			i := int(c - 'a')
			if !qs[i] {
				count++
				qs[i] = true
			}
		}
	}
	return count
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
