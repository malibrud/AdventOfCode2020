package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	// Parse the jolt adapters
	var adapters []int
	N := len(lines)
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		adapters = append(adapters, v)
	}

	sort.Ints(adapters)

	var diffs []int
	count1 := 0
	count2 := 0
	count3 := 0
	lastAdapter := 0
	var i int
	for i = 0; i < N; i++ {
		diff := adapters[i] - lastAdapter
		diffs = append(diffs, diff)
		switch diff {
		case 1:
			count1++
		case 2:
			count2++
		case 3:
			count3++
		}
		lastAdapter = adapters[i]
	}
	diffs = append(diffs, 3)
	N = len(diffs)

	// observation: There are no differences of 2 in the example data or in the provided data
	// Need to look for runs of 1's between the 3's.
	// There is only one way to make a jump of 3.

	prod := int64(1)
	runLen := 0 // number of 1's in a row
	for i = 0; i < N; i++ {
		switch diffs[i] {
		case 1:
			runLen++
		case 2:
			panic("Error: Differences of 2 are not supported")
		case 3:
			switch runLen {
			case 0:
				// only 1 way a run of 0 can happen
				// {}
				prod *= 1
			case 1:
				// only 1 way a run of 1 can happen
				// {1}
				prod *= 1
			case 2:
				// only 2 ways a run of 2 can happen
				// {1,1}, {2}
				prod *= 2
			case 3:
				// only 4 ways a run of 3 can happen
				// {1,1,1}, {1,2}, {2,1}, {3}
				prod *= 4
			case 4:
				// only 7 ways a run of 4 can happen
				// {1,1,1,}
				// {1,1,2}, {1,2,1}, {2,1,1}
				// {2,2}
				// {1,3}, {3,1}
				prod *= 7
			default:
				panic("Error 1-runs of length > 4 are not supported")
			}
			runLen = 0
		default:
			panic("Error: Differences > 3 are not supported")
		}
	}
	return strconv.FormatInt(prod, 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
