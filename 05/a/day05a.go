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

	maxSeat := 0
	for _, bp := range boardingPasses {
		num := bp
		num = strings.Replace(num, "F", "0", -1)
		num = strings.Replace(num, "B", "1", -1)
		num = strings.Replace(num, "L", "0", -1)
		num = strings.Replace(num, "R", "1", -1)

		seat, _ := strconv.ParseInt(num, 2, 32)
		if int(seat) > maxSeat {
			maxSeat = int(seat)
		}
	}
	return strconv.Itoa(maxSeat)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
