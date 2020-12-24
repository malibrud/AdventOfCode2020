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
	G := 6
	R0 := len(lines)
	C0 := len(lines[0])
	R := R0 + 2*G
	C := C0 + 2*G
	H := 1 + 2*G
	W := 1 + 2*G

	// Create two buffers that will toggle back and forth
	var b, r, c, h, w int
	var cubes [2][][][][]bool
	for b = 0; b < 2; b++ {
		cubes[b] = make([][][][]bool, R)
		for r = 0; r < R; r++ {
			cubes[b][r] = make([][][]bool, C)
			for c := 0; c < C; c++ {
				cubes[b][r][c] = make([][]bool, H)
				for h := 0; h < H; h++ {
					cubes[b][r][c][h] = make([]bool, W)
				}
			}
		}
	}

	// Internalize the input data
	for r = 0; r < R0; r++ {
		for c := 0; c < C0; c++ {
			cubes[0][r+G][c+G][G][G] = (lines[r][c] == '#')
		}
	}

	// Array of the 80 differential positions
	// to automate the checking of adjacent cubes
	N := 80
	var diffs [80][4]int
	i := 0
	for r = -1; r < 2; r++ {
		for c = -1; c < 2; c++ {
			for h = -1; h < 2; h++ {
				for w = -1; w < 2; w++ {
					if r != 0 || c != 0 || h != 0 || w != 0 {
						diffs[i][0] = r
						diffs[i][1] = c
						diffs[i][2] = h
						diffs[i][3] = w
						i++
					}
				}
			}
		}
	}

	// Simulate the Cellular automaton
	s := 0
	d := 1
	count := 0
	for g := 0; g < G; g++ {
		count = 0
		for r := 0; r < R; r++ {
			for c := 0; c < C; c++ {
				for h = 0; h < H; h++ {
					for w = 0; w < W; w++ {
						cube := cubes[s][r][c][h][w]
						// Check the 8 neighbors and count occupied positions
						neighbors := 0
						for n := 0; n < N; n++ {
							// walk the dog to all nieghboring cubes using diffs
							ir := r + diffs[n][0]
							ic := c + diffs[n][1]
							ih := h + diffs[n][2]
							iw := w + diffs[n][3]
							if ir < 0 || ir >= R || ic < 0 || ic >= C || ih < 0 || ih >= H || iw < 0 || iw >= W {
								continue
							}
							if cubes[s][ir][ic][ih][iw] {
								neighbors++
							}
						}
						if cube && (neighbors < 2 || neighbors > 3) {
							// Active --> Inactive
							cubes[d][r][c][h][w] = false
						} else if !cube && neighbors == 3 {
							// Inactive --> Active
							cubes[d][r][c][h][w] = true
						} else {
							// No change
							cubes[d][r][c][h][w] = cube
						}
						if cubes[d][r][c][h][w] {
							count++
						}
					}
				}
			}
		}

		println("Gen: ", g, "  Count:", count)

		// Swap source and destination buffers
		d, s = s, d
	}

	return strconv.Itoa(count)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
