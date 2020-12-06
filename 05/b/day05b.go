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
	boardingPasses := strings.Split(string(dat), "\r\n")

	var seats [128][8]int

	for _, bp := range boardingPasses {
		num := bp
		num = strings.Replace(num, "F", "0", -1)
		num = strings.Replace(num, "B", "1", -1)
		num = strings.Replace(num, "L", "0", -1)
		num = strings.Replace(num, "R", "1", -1)

		seat, _ := strconv.ParseInt(num, 2, 32)
		r := seat >> 3
		c := seat & 7
		seats[r][c] = 1
	}
	plot(&seats)

	s := findMySeat(&seats)
	return strconv.Itoa(s)
}

func plot(s *[128][8]int) {
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			if s[i][j] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
}

func findMySeat(s *[128][8]int) int {
	var seat int
	nSeats := 128 * 8
	for seat = 0; seat < nSeats; seat++ {
		r := seat >> 3
		c := seat & 7
		if s[r][c] == 1 {
			break
		}
	}
	for ; seat < nSeats; seat++ {
		r := seat >> 3
		c := seat & 7
		if s[r][c] == 0 {
			break
		}
	}
	return seat
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
