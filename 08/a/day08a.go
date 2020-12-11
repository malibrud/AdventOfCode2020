package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	op      string
	arg     int
	visited bool
}

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	var program []instruction

	// Parse the program
	for _, line := range lines {
		parts := strings.Split(line, " ")
		op := parts[0]
		arg, _ := strconv.Atoi(parts[1])
		i := instruction{op, arg, false}
		program = append(program, i)
	}

	// Execute the program
	acc := 0
	ic := 0
	for {
		if program[ic].visited {
			break
		}
		program[ic].visited = true

		switch program[ic].op {
		case "nop":
			ic++
		case "acc":
			acc += program[ic].arg
			ic++
		case "jmp":
			ic += program[ic].arg
		}
	}

	return strconv.Itoa(acc)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
