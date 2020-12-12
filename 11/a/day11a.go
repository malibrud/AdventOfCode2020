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
	R := len(lines)
	C := len(lines[0])

	// Parse the seats
	// Create two buffers that will toggle back and forth
	var seats [2][][]byte
	seats[0] = make([][]byte, R)
	seats[1] = make([][]byte, R)
	for r := 0; r < R; r++ {
		seats[0][r] = make([]byte, C)
		seats[1][r] = make([]byte, C)
		for c := 0; c < C; c++ {
			seats[0][r][c] = lines[r][c]
			seats[1][r][c] = lines[r][c]
		}
	}

	// Array of the 8 differential positions
	// to automate the checking of adjacent seats
	diffs := [8][2]int{
		// {r, c}
		// row ahead
		{-1, -1},
		{-1, 0},
		{-1, +1},
		// this row
		{0, -1},
		{0, +1},
		// row behind
		{+1, -1},
		{+1, 0},
		{+1, +1},
	}
	N := 8

	// Simulate the Cellular automaton
	iteration := 1
	s := 0
	d := 1
	for {
		same := true // if this stays true, no seats changed
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				seat := seats[s][r][c]
				// Check to see if this position is on the floor
				if seat == '.' {
					continue
				}
				// Check the 8 neighbors and count occupied positions
				neighbors := 0
				for n := 0; n < N; n++ {
					// walk the dog to all nieghboring seats using diffs
					ir := r + diffs[n][0]
					ic := c + diffs[n][1]
					// Check to see if we are off the edge
					if ir < 0 || ir >= R || ic < 0 || ic >= C {
						continue
					}
					if seats[s][ir][ic] == '#' {
						neighbors++
					}
				}
				if seat == 'L' && neighbors == 0 {
					// Plenty of room, occupy the seat
					seats[d][r][c] = '#'
					same = false
				} else if seat == '#' && neighbors >= 4 {
					// Too crowded, vacate the seat
					seats[d][r][c] = 'L'
					same = false
				} else {
					// Seat OK.  Leave it alone
					seats[d][r][c] = seat // not necessary
				}
			}
		}
		// Break out if the grids are the same
		if same {
			break
		}
		// Swap source and destination buffers
		d, s = s, d

		iteration++
	}

	// Finally count the occupied seats
	count := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if seats[d][r][c] == '#' {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
