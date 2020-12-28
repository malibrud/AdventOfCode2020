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

func peekOperator(expr string, pcin int) (byte, int) {
	if pcin+1 >= len(expr) {
		return byte(0), 0
	}
	op := expr[pcin+1]
	if op == ')' {
		return byte(0), 0
	}

	if op == '+' || op == '*' {
		return op, pcin + 1
	}

	// recurse through parens
	depth := 0
	var pc int
	for pc = pcin; pc < len(expr); pc++ {
		if expr[pc] == '(' {
			depth++
		}
		if expr[pc] == ')' {
			depth--
		}
		if depth == 0 {
			break
		}
	}
	if pc+1 >= len(expr) || expr[pc+1] == ')' {
		return byte(0), 0
	}
	return expr[pc+1], pc + 1
}

func evalOperand(expr string, pcin int) (result int, pcout int) {
	pc := pcin
	val := expr[pc]
	lhs := int(val - '0')
	var rhs int
	opn, pcon := peekOperator(expr, pc)

	// Get LHS
	if val == '(' {
		lhs, pc = eval(expr, pc+1)
		if expr[pc] != ')' {
			panic("Expected ')'")
		}
	} else if '0' <= val && val <= '9' {
		result = lhs
	} else {
		panic("parse error")
	}
	pc++

	if opn == '*' || opn == byte(0) {
		result = lhs
	} else if opn == '+' {
		rhs, pc = evalOperand(expr, pcon+1)
		result = lhs + rhs
	}
	return result, pc
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
