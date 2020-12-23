package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
	name       string
	r1lb, r1ub int // lower and upper bound for range 1
	r2lb, r2ub int // lower and upper bound for range 2
}

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName string, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	blocks := strings.Split(string(dat), "\r\n\r\n")

	// Parse Fields
	re := regexp.MustCompile(`([a-z]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	var fields []field
	lines := strings.Split(blocks[0], "\r\n")
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		var f field
		f.name = matches[1]
		f.r1lb, _ = strconv.Atoi(matches[2])
		f.r1ub, _ = strconv.Atoi(matches[3])
		f.r2lb, _ = strconv.Atoi(matches[4])
		f.r2ub, _ = strconv.Atoi(matches[5])
		fields = append(fields, f)
	}

	// Parse my ticket
	// Skip for part 1

	// Parse nearby tickets
	sum := 0
	lines = strings.Split(blocks[2], "\r\n")[1:]
	for _, line := range lines {
		vals := strings.Split(line, ",")
		for _, val := range vals {
			v, _ := strconv.Atoi(val)
			found := false
			for _, f := range fields {
				if (f.r1lb <= v && v <= f.r1ub) || (f.r2lb <= v && v <= f.r2ub) {
					found = true
				}
			}
			if !found {
				sum += v
			}
		}
	}

	return strconv.FormatInt(int64(sum), 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
