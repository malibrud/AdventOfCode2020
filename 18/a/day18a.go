package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")

	sum := 0
	for _, line := range lines {
		l := strings.ReplaceAll(line, " ", "")
		result, _ := eval(l, 0)
		sum += result
	}

	return strconv.Itoa(sum)
}

func eval(expr string, pcin int) (result int, pcout int) {
	var lhs, rhs int
	pc := pcin

	// Get LHS
	lhs, pc = evalOperand(expr, pc)

	for pc < len(expr) && expr[pc] != ')' {
		// Parse operator
		op := expr[pc]
		pc++

		// Parse the RHS
		rhs, pc = evalOperand(expr, pc)

		// compute the result
		switch op {
		case '+':
			lhs += rhs
		case '*':
			lhs *= rhs
		}
	}
	return lhs, pc
}

func evalOperand(expr string, pcin int) (result int, pcout int) {
	pc := pcin
	val := expr[pc]
	if val == '(' {
		result, pc = eval(expr, pc+1)
		if expr[pc] != ')' {
			panic("Expected ')'")
		}
		pc++
	} else if '0' <= val && val <= '9' {
		result = int(val - '0')
		pc++
	} else {
		panic("parse error")
	}
	return result, pc
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
