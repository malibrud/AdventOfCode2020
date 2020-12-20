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

func doit(fileName string, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fields := strings.Split(string(dat), ",")
	turns := make(map[int]int)

	// Initialize the game
	LAST_TURN := 30_000_000
	turn := 2
	v, _ := strconv.Atoi(fields[0])
	lastSpoken := v
	for i := 1 ; i < len(fields) ; i++ {
		v, _ := strconv.Atoi(fields[i])
		turns[lastSpoken] = turn - 1 
		lastSpoken = v
		turn++
	}

	// Play the game
	for ; turn <= LAST_TURN; turn++ {
		var v int
		if _, ok := turns[lastSpoken];ok {
			v = (turn - 1) - turns[lastSpoken]
		} else {
			v = 0
		}
		turns[lastSpoken] = turn - 1
		lastSpoken = v
	}

	return strconv.Itoa(lastSpoken)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
