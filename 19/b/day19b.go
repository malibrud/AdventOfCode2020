package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule struct {
	val  byte
	runs [][]int
}

type Matcher struct {
	rules        map[int]*Rule
	depth        int
	maxRule8Runs int
}

/// NewMatcher - Construct a new matcher.
func NewMatcher() *Matcher {
	m := new(Matcher)
	m.rules = make(map[int]*Rule)
	m.depth = 0
	m.maxRule8Runs = 0
	return m
}

func (m Matcher) addRule(ruleLine string) {
	rule := new(Rule)
	parts := strings.Split(ruleLine, ": ")
	if parts[0] == "8" {
		// replace rule 8 with "8: 42 | 42 8"
		// Swap the order to make it greedy
		parts[1] = "42 8 | 42"
	}
	if parts[0] == "11" {
		// replace rule 11 with "11: 42 31 | 42 11 31"
		// Swap the order to make it greedy
		parts[1] = "42 11 31 | 42 31"
	}
	n, _ := strconv.Atoi(parts[0])
	if parts[1][0] == '"' {
		rule.val = parts[1][1]
		rule.runs = nil
	} else {
		rule.val = byte(0)
		runs := strings.Split(parts[1], " | ")
		rule.runs = make([][]int, len(runs))
		for i := 0; i < len(runs); i++ {
			vals := strings.Split(runs[i], " ")
			rule.runs[i] = make([]int, len(vals))
			for j := 0; j < len(vals); j++ {
				val, _ := strconv.Atoi(vals[j])
				rule.runs[i][j] = val
			}
		}
	}
	m.rules[n] = rule
}

func (m Matcher) match(ir int, message string, b int) (bool, int) {
	r := m.rules[ir]
	pad := strings.Repeat("  ", m.depth)
	fmt.Printf("%sTrying Rule %d\n", pad, ir)
	if b >= len(message) {
		return false, 0
	}
	if r.val != 0 && message[b] == r.val {
		return true, 1
	}
	for i := 0; i < len(r.runs); i++ {
		// Try each run
		run := r.runs[i]
		p := b
		success := true
		for j := 0; j < len(run); j++ {
			m.depth++
			OK, l := m.match(run[j], message, p)
			m.depth--
			success = success && OK
			if !success {
				break
			}
			p += l
		}
		if success {
			fmt.Printf("%sRule %d, %d: Matched [%d, %d] %s\n", pad, ir, i, b, p, message[b:p])
			fmt.Printf("%s    %s\n", pad, message)
			leadingSpaces := strings.Repeat(" ", b)
			middleSpaces := strings.Repeat(" ", p-b-1)
			fmt.Printf("%s    %s^%sx\n", pad, leadingSpaces, middleSpaces)

			return true, p - b
		}
	}
	return false, 0
}

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	parts := strings.Split(string(dat), "\r\n\r\n")
	rulePart := parts[0]
	dataPart := parts[1]

	// parse the rules
	matcher := NewMatcher()
	for _, ruleLine := range strings.Split(rulePart, "\r\n") {
		matcher.addRule(ruleLine)
	}

	count := 0
	for _, dataLine := range strings.Split(dataPart, "\r\n") {
		fmt.Printf("\n\nTesting: %s\n", dataLine)
		for i := 1; i < len(dataLine)-1; i++ {
			// Very inefficient approach... should be improved.
			front := dataLine[0:i]
			back := dataLine[i:]
			OKf, lf := matcher.match(8, front, 0)
			OKb, lb := matcher.match(11, back, 0)
			if OKf && OKb && lf == len(front) && lb == len(back) {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
