package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	h := len(lines)
	w := len(lines[0])

	slopes := []coord{
		{x: 1, y: 1},
		{x: 3, y: 1},
		{x: 5, y: 1},
		{x: 7, y: 1},
		{x: 1, y: 2},
	}

	hitProduct := 1
	for _, slope := range slopes {
		dx := slope.x
		dy := slope.y
		pos := 0
		hitCount := 0
		for i := 0; i < h; i += dy {
			line := lines[i]
			if line[pos] == '#' {
				hitCount++
			}
			pos = (pos + dx) % w
		}
		hitProduct *= hitCount
	}
	return strconv.Itoa(hitProduct)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
