package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bagSet struct {
	count int
	color string
}

type bagContents struct {
	contents *[]bagSet
	found    bool
}

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	bagLines := strings.Split(string(dat), "\r\n")
	//nBags := len(bagLines)
	bags := make(map[string]*bagContents)

	for _, bagLine := range bagLines {
		L := len(bagLine)
		pair := strings.Split(bagLine[:L-1], " contain ")
		L = len(pair[0])
		color := pair[0][:L-5]
		list := pair[1]
		if list == "no other bags" {
			bags[color] = &bagContents{nil, false}
		} else {
			contents := strings.Split(list, ", ")
			var a []bagSet
			for _, set := range contents {
				fields := strings.Split(set, " ")
				n, _ := strconv.Atoi(fields[0])
				c := fields[1] + " " + fields[2]
				b := bagSet{n, c}
				a = append(a, b)
			}
			bags[color] = &bagContents{&a, false}
		}
	}

	for k := range bags {
		markBagsContainingShinyGold(bags, k)
	}

	total := 0
	for k := range bags {
		if bags[k].found {
			total++
		}
	}

	return strconv.Itoa(total)
}

func markBagsContainingShinyGold(bags map[string]*bagContents, color string) bool {
	if color == "shiny gold" {
		return true
	}
	if bags[color].found {
		return true
	}
	if bags[color].contents == nil {
		return false
	}
	found := false
	for _, b := range *bags[color].contents {
		found = found || markBagsContainingShinyGold(bags, b.color)
	}
	bags[color].found = found
	return found
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
