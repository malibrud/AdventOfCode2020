package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt", 25))
}

func doit(fileName string, preamble int) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")

	// Parse the cipher
	var xmasData []int64
	N := len(lines)
	for _, line := range lines {
		v, _ := strconv.ParseInt(line, 10, 64)
		xmasData = append(xmasData, v)
	}

	var i int
	for i = preamble; i < N; i++ {
		match := false
	searchLoop:
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if xmasData[j]+xmasData[k] == xmasData[i] {
					match = true
					break searchLoop
				}
			}
		}
		if !match {
			break
		}
	}
	return strconv.FormatInt(xmasData[i], 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
