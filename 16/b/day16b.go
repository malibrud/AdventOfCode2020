package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type field struct {
	name       string
	r1lb, r1ub int // lower and upper bound for range 1
	r2lb, r2ub int // lower and upper bound for range 2
	idx        int
	count      int
}

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName string, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	blocks := strings.Split(string(dat), "\r\n\r\n")

	// Parse Fields
	re := regexp.MustCompile(`([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
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
		f.idx = -1
		fields = append(fields, f)
	}
	nFields := len(fields)

	// Parse my ticket
	var myticket []int
	lines = strings.Split(blocks[1], "\r\n")
	vals := strings.Split(lines[1], ",")
	for _, val := range vals {
		v, _ := strconv.Atoi(val)
		myticket = append(myticket, v)
	}

	// Parse nearby tickets, remembering valid ones.
	var validTickets [][]int
	lines = strings.Split(blocks[2], "\r\n")[1:]
	for _, line := range lines {
		vals := strings.Split(line, ",")
		t := make([]int, nFields)
		valid := true
		for j, val := range vals {
			v, _ := strconv.Atoi(val)
			t[j] = v
			foundField := false
			for _, f := range fields {
				if (f.r1lb <= v && v <= f.r1ub) || (f.r2lb <= v && v <= f.r2ub) {
					foundField = true
				}
			}
			if !foundField {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}

	// Find comonality count
	for i, f := range fields { // Field definitions
		count := 0
		for j := range fields { // nearby ticket field
			match := true
			for k := range validTickets { // nearby tickets
				v := validTickets[k][j]
				if !(f.r1lb <= v && v <= f.r1ub) && !(f.r2lb <= v && v <= f.r2ub) {
					match = false
					break
				}
			}
			if match {
				count++
			}
		}
		fields[i].count = count
	}

	// Sort by count
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].count < fields[j].count
	})

	// Rerun the matching process.
	visited := make([]bool, nFields)
	for i, f := range fields { // Field definitions
		for j := range fields { // nearby ticket field
			match := true
			for k := range validTickets { // nearby tickets
				v := validTickets[k][j]
				if !(f.r1lb <= v && v <= f.r1ub) && !(f.r2lb <= v && v <= f.r2ub) {
					match = false
					break
				}
			}
			if match && !visited[j] {
				fields[i].idx = j
				visited[j] = true
				break
			}
		}
	}

	// Multiply my ticket fields whith "departure"
	prod := int64(1)
	for _, f := range fields {
		if strings.Contains(f.name, "departure") {
			prod *= int64(myticket[f.idx])
		}
	}

	return strconv.FormatInt(int64(prod), 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
