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

	pos := 0
	hitCount := 0
	w := len(lines[0])
	for _, line := range lines {
		if line[pos] == '#' {
			hitCount++
		}
		pos = (pos + 3) % w
	}
	return strconv.Itoa(hitCount)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
