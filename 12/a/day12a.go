package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt"))
}

func doit(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")

	// Start at (0,0)
	px := 0 // E-W coordinate, + = E
	py := 0 // N-S coordinate, + = N
	// Start heading E
	hx := 1 // E-W heading, + = E
	hy := 0 // N-S heading, + = N
	for _, line := range lines {
		action := line[0]
		arg, _ := strconv.Atoi(line[1:])
		switch action {
		case 'N':
			py += arg
		case 'S':
			py -= arg
		case 'E':
			px += arg
		case 'W':
			px -= arg
		case 'L':
			re, im := getC(arg)
			hx, hy = hx*re-hy*im, hx*im+hy*re
		case 'R':
			re, im := getC(-arg)
			hx, hy = hx*re-hy*im, hx*im+hy*re
		case 'F':
			px += arg * hx
			py += arg * hy
		default:
			panic("instruction not found")
		}
	}

	d := abs(px) + abs(py)
	return strconv.Itoa(d)
}

func getC(ang int) (re, im int) {
	if ang < 0 {
		ang += 360
	}
	switch ang {
	case 0:
		return 1, 0
	case 90:
		return 0, 1
	case 180:
		return -1, 0
	case 270:
		return 0, -1
	default:
		panic("Angle not found")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
