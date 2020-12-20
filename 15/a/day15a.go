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
	var numbers [2021]int

	// Initialize the game
	turn := 1
	for _, field := range fields {
		v, _ := strconv.Atoi(field)
		numbers[turn] = v
		turn++
	}

	// Play the game
	for ; turn <= 2020; turn++ {
		lastSpoken := numbers[turn-1]
		found := false
		for lb := turn - 2; lb >= 1; lb-- {
			if numbers[lb] == lastSpoken {
				// Number was found
				numbers[turn] = (turn - 1) - lb
				found = true
				break
			}
		}
		if !found {
			// Number was not found
			numbers[turn] = 0
		}
	}

	return strconv.Itoa(numbers[2020])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
