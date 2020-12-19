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

	earliestTime, _ := strconv.Atoi(lines[0])
	buses := strings.Split(lines[1], ",")
	var ids []int

	for _, v := range buses {
		if v != "x" {
			t, _ := strconv.Atoi(v)
			ids = append(ids, t)
		}
	}

	min := 1000000
	id := 0
	for _, v := range ids {
		dist := v - (earliestTime % v)
		if dist < min {
			min = dist
			id = v
		}
	}
	ans := id * min
	return strconv.Itoa(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
