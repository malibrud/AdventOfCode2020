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
	rules map[int]*Rule
}

/// NewMatcher - Construct a new matcher.
func NewMatcher() *Matcher {
	m := new(Matcher)
	m.rules = make(map[int]*Rule)
	return m
}

func (m Matcher) addRule(ruleLine string) {
	rule := new(Rule)
	parts := strings.Split(ruleLine, ": ")
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

func (m Matcher) match(ir int, message string, b int, e int) (bool, int) {
	r := m.rules[ir]
	if r.val != 0 && message[b] == r.val {
		return true, 1
	}
	for i := 0; i < len(r.runs); i++ {
		// Try each run
		run := r.runs[i]
		p := b
		success := true
		for j := 0; j < len(run); j++ {
			OK, l := m.match(run[j], message, p, e)
			success = success && OK
			if !success {
				break
			}
			p += l
		}
		if success {
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

		if OK, l := matcher.match(0, dataLine, 0, len(dataLine)); OK && l == len(dataLine) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
