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
	// Append end instruction to simplify detection
	i := instruction{"end", 0, false}
	program = append(program, i)
	N := len(program)

	swap := map[string]string{"nop": "jmp", "jmp": "nop", "acc": "acc"}
	// Execute the program
	acc := 0
	success := false
	for i := 0; i < N; i++ {
		program[i].op = swap[program[i].op]
		acc, success = runProgram(program)
		program[i].op = swap[program[i].op]
		if success {
			break
		}
		// reset visited flags
		for j := 0; j < N; j++ {
			program[j].visited = false
		}
	}

	return strconv.Itoa(acc)
}

func runProgram(program []instruction) (int, bool) {
	acc := 0
	ic := 0
	for {
		if program[ic].visited {
			return acc, false
		}
		if program[ic].op == "end" {
			return acc, true
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
	return 0, false // Should not be reachable
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
