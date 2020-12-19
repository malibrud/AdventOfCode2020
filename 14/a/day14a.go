package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName string, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	mem := make(map[int]int64)

	dcmask := int64(0)
	setmask := int64(0)
	re := regexp.MustCompile(`mem\[(\d+)\]`)
	for _, line := range lines {
		instruct := strings.Split(line, " = ")
		lval := instruct[0]
		rval := instruct[1]
		if lval == "mask" {
			var dcstr, setstr string

			// Mask of don't care values, 1 = 0 leave it alone
			dcstr = strings.ReplaceAll(rval, "1", "0")
			dcstr = strings.ReplaceAll(rval, "X", "1")
			dcmask, _ = strconv.ParseInt(dcstr, 2, 64)

			// Mask of don't care values, 1 = 0 leave it alone
			setstr = strings.ReplaceAll(rval, "X", "0")
			setmask, _ = strconv.ParseInt(setstr, 2, 64)
		} else if lval[:3] == "mem" {
			fields := re.FindStringSubmatch(lval)
			addr, _ := strconv.Atoi(fields[1])
			val, _ := strconv.ParseInt(rval, 10, 64)

			// Do the masking thing
			result := val&dcmask | setmask
			mem[addr] = result
		}
	}
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}

	return strconv.FormatInt(sum, 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
