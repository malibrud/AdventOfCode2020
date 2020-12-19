package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bus struct {
	id    int64
	place int64
}

func main() {
	fmt.Println(doit("data.txt", "100000000000000"))
}

func doit(fileName string, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	startIdx, _ := strconv.ParseInt(arg, 10, 64)

	fields := strings.Split(lines[1], ",")

	var buses []bus

	for i, v := range fields {
		if v != "x" {
			t, _ := strconv.Atoi(v)
			buses = append(buses, bus{int64(t), int64(i)})
		}
	}
	N := len(buses)

	step := int64(1)
	v := startIdx
	for j := 0; j < N; j++ {
		id := buses[j].id
		place := buses[j].place % id // It may be that place > id
		for {
			if v%id == (id-place)%id {
				step *= id
				println("Found, id =", id, "place =", place, "v =", v)
				break
			}
			v += step
		}
	}

	return strconv.FormatInt(v, 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
