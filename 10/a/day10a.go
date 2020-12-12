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

	lastAdapter := 0
	count1 := 0
	count3 := 1 // includes the adapter not in the list
	for i := 0; i < N; i++ {
		diff := adapters[i] - lastAdapter
		lastAdapter = adapters[i]
		switch diff {
		case 1:
			count1++
		case 3:
			count3++
		}
	}
	ans := count1 * count3
	return strconv.Itoa(ans)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
