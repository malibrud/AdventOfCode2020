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

	// Find index which meets criteria for part a
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

	targetVal := xmasData[i]

	var max, min, sum int64
sumLoop:
	for s := 0; s < N; s++ {
		max = xmasData[s]
		min = xmasData[s]
		sum = xmasData[s]
		for i = s + 1; i < N && sum <= targetVal; i++ {
			v := xmasData[i]
			max = getmax(max, v)
			min = getmin(min, v)
			sum += v
			if sum == targetVal {
				break sumLoop
			}
		}
	}
	total := max + min
	return strconv.FormatInt(total, 10)
}

func getmax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func getmin(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
