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
	// Start waypoint initialized to (10,1)
	wx := 10 // E-W waypoint, + = E
	wy := 1  // N-S waypoint, + = N
	for _, line := range lines {
		action := line[0]
		arg, _ := strconv.Atoi(line[1:])
		switch action {
		case 'N':
			wy += arg
		case 'S':
			wy -= arg
		case 'E':
			wx += arg
		case 'W':
			wx -= arg
		case 'L':
			re, im := getC(arg)
			wx, wy = wx*re-wy*im, wx*im+wy*re
		case 'R':
			re, im := getC(-arg)
			wx, wy = wx*re-wy*im, wx*im+wy*re
		case 'F':
			px += arg * wx
			py += arg * wy
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
