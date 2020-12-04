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
	nums := make([]int, 0, len(lines))

	for _, s := range lines {
		n, e := strconv.Atoi(s)
		check(e)
		nums = append(nums, n)
	}

	for _, n1 := range nums {
		for _, n2 := range nums {
			if (n1 + n2) == 2020 {
				return strconv.Itoa(n1 * n2)
			}
		}
	}
	return "Error"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
