package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type tileTy struct {
	id        int
	data      []string
	H, W, N   int
	topEdge   bool
	botEdge   bool
	leftEdge  bool
	rightEdge bool
	top       string
	bot       string
	left      string
	right     string
	topR      string
	botR      string
	leftR     string
	rightR    string
}

func (t *tileTy) print() {
	fmt.Println("ID:", t.id)
	for i := 0; i < t.N; i++ {
		fmt.Println(t.data[i])
	}
}

func (t *tileTy) rotate(R []int) {
	re := R[0]
	im := R[1]
	D := t.N - 1

	d := make([][]byte, t.N)
	for i := 0; i < t.N; i++ {
		d[i] = make([]byte, t.N)
	}
	for x := 0; x < t.N; x++ {
		for y := 0; y < t.N; y++ {
			xn, yn := 2*x-D, 2*y-D
			xn, yn = xn*re-yn*im, xn*im+yn*re
			xn, yn = (xn+D)/2, (yn+D)/2
			d[xn][yn] = t.data[x][y]
		}
	}
	for x := 0; x < t.N; x++ {
		t.data[x] = string(d[x])
	}

	t.internalizeXEdges()
}

func (t *tileTy) xformEdgeToLeft(e string) {
	R := make([]int, 2)
	if e == t.left || e == t.leftR {
		R[0], R[1] = 1, 0
	} else if e == t.top || e == t.topR {
		R[0], R[1] = 0, 1
	} else if e == t.right || e == t.rightR {
		R[0], R[1] = -1, 0
	} else if e == t.bot || e == t.botR {
		R[0], R[1] = 0, -1
	} else {
		panic("Edge not found")
	}
	t.rotate(R)
}

func (t *tileTy) hasEdge(e string) bool {
	if e == t.top || e == t.topR ||
		e == t.bot || e == t.botR ||
		e == t.left || e == t.leftR ||
		e == t.right || e == t.rightR {
		return true
	}
	return false
}

func (t *tileTy) xformEdgeToTop(e string) {
	R := make([]int, 2)
	if e == t.top || e == t.topR {
		R[0], R[1] = 1, 0
	} else if e == t.right || e == t.rightR {
		R[0], R[1] = 0, 1
	} else if e == t.bot || e == t.botR {
		R[0], R[1] = -1, 0
	} else if e == t.left || e == t.leftR {
		R[0], R[1] = 0, -1
	} else {
		panic("Edge not found")
	}
	t.rotate(R)
}

func (t *tileTy) flipToMatchOnLeft(e string) {
	if e == t.leftR {
		for i := 0; i < t.N/2; i++ {
			j := t.N - 1 - i
			t.data[i], t.data[j] = t.data[j], t.data[i]
		}
	}
	t.internalizeXEdges()
}

func (t *tileTy) flipToMatchOnTop(e string) {
	if e == t.topR {
		for i := 0; i < t.N; i++ {
			t.data[i] = reverse(t.data[i])
		}
	}
	t.internalizeXEdges()
}

func (t *tileTy) internalizeXEdges() {
	t.top = string(t.data[0])
	t.bot = string(t.data[t.H-1])
	t.left = getCol(t.data, 0)
	t.right = getCol(t.data, t.W-1)
	t.topR = reverse(t.top)
	t.botR = reverse(t.bot)
	t.leftR = reverse(t.left)
	t.rightR = reverse(t.right)
}

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	rawTiles := strings.Split(string(dat), "\r\n\r\n")

	fmt.Printf("Number of tiles found: %d\n", len(rawTiles))
	G := int(math.Sqrt(float64(len(rawTiles))))
	edgeCounts := make(map[string]int)
	tileIds := make(map[string]int)
	tileMap := make(map[int]*tileTy)
	idSet := make(map[int]bool)

	// parse the tiles
	for _, rawTile := range rawTiles {
		lines := strings.Split(rawTile, "\r\n")
		idstr := lines[0][5 : len(lines[0])-1]
		id, _ := strconv.Atoi(idstr)
		tile := lines[1:len(lines)]
		t := new(tileTy)
		t.H = len(tile)
		t.W = len(tile[0])
		t.N = t.W
		t.id = id
		t.data = tile
		t.top = string(tile[0])
		t.bot = string(tile[t.H-1])
		t.left = getCol(tile, 0)
		t.right = getCol(tile, t.W-1)
		t.topR = reverse(t.top)
		t.botR = reverse(t.bot)
		t.leftR = reverse(t.left)
		t.rightR = reverse(t.right)
		tileMap[id] = t
		idSet[id] = true
		edgeCounts[t.top]++
		edgeCounts[t.bot]++
		edgeCounts[t.left]++
		edgeCounts[t.right]++
		edgeCounts[t.topR]++
		edgeCounts[t.botR]++
		edgeCounts[t.leftR]++
		edgeCounts[t.rightR]++
		tileIds[t.top] = id
		tileIds[t.bot] = id
		tileIds[t.left] = id
		tileIds[t.right] = id
		tileIds[t.topR] = id
		tileIds[t.botR] = id
		tileIds[t.leftR] = id
		tileIds[t.rightR] = id
	}

	cornerTiles := make(map[int]int)
	for k, v := range edgeCounts {
		if v == 1 {
			id := tileIds[k]
			cornerTiles[id]++
			t := tileMap[id]
			if k == t.top || k == t.topR {
				t.topEdge = true
			} else if k == t.bot || k == t.botR {
				t.botEdge = true
			} else if k == t.left || k == t.leftR {
				t.leftEdge = true
			} else if k == t.right || k == t.rightR {
				t.rightEdge = true
			}
		}
	}

	prod := int64(1)
	cornerIDs := make([]int, 4)
	i := 0
	for id, count := range cornerTiles {
		if count == 4 {
			prod *= int64(id)
			cornerIDs[i] = id
			i++
		}
	}

	// Assume GxG grid based on 144 tiles
	grid := make([][]int, G)
	for i := 0; i < G; i++ {
		grid[i] = make([]int, G)
	}
	// Choose a corner to start with
	grid[0][0] = cornerIDs[0]
	t := tileMap[grid[0][0]]
	delete(idSet, grid[0][0])
	R := make([]int, 2)
	if t.leftEdge && t.topEdge {
		R[0] = 1
		R[1] = 0
	}
	if t.topEdge && t.rightEdge {
		R[0] = 0
		R[1] = 1
	}
	if t.rightEdge && t.botEdge {
		R[0] = -1
		R[1] = 0
	}
	if t.botEdge && t.leftEdge {
		R[0] = 0
		R[1] = -1
	}
	t.rotate(R)
	t.print()

	// Solve the top row
	for c := 1; c < G; c++ {
		lt := tileMap[grid[0][c-1]]
		re := lt.right
		id := -1
		for k := range idSet {
			if tileMap[k].hasEdge(re) {
				id = k
				break
			}
		}
		if id == -1 {
			panic("Tile not found for edge")
		}
		t := tileMap[id]
		t.xformEdgeToLeft(re)
		t.flipToMatchOnLeft(re)
		grid[0][c] = id
		delete(idSet, id)
		t.print()
	}

	// Solve the remaining rows
	for r := 1; r < G; r++ {
		for c := 0; c < G; c++ {
			tt := tileMap[grid[r-1][c]]
			be := tt.bot
			id := -1
			for k := range idSet {
				if tileMap[k].hasEdge(be) {
					id = k
					break
				}
			}
			if id == -1 {
				panic("Tile not found for edge")
			}
			t := tileMap[id]
			t.xformEdgeToTop(be)
			t.flipToMatchOnTop(be)
			grid[r][c] = id
			delete(idSet, id)
			t.print()
		}
	}

	N := tileMap[grid[0][0]].N
	mapRows := G * (N - 2)
	b := 1
	e := N - 1
	sep := ""
	sepRows := 0
	//mapRows := G * (N - 0)
	//b := 0
	//e := N
	//sep := " "
	//sepRows := 1
	compMap := make([]string, mapRows)
	i = 0
	for rg := 0; rg < G; rg++ {
		for tg := b; tg < e; tg++ {
			line := ""
			for cg := 0; cg < G; cg++ {
				t := tileMap[grid[rg][cg]]
				line += t.data[tg][b:e] + sep
			}
			compMap[i] = line
			i++
			fmt.Println(line)
		}
		for j := 0; j < sepRows; j++ {
			fmt.Println("")
		}
	}
	D := mapRows

	dat, err = ioutil.ReadFile("monster.txt")
	check(err)
	nHashes := strings.Count(string(dat), "#")
	monster := strings.Split(string(dat), "\r\n")
	MW := len(monster[0])
	MH := len(monster)
	monsterCoords := make([][]int, nHashes)
	im := 0
	for i := 0; i < MH; i++ {
		for j := 0; j < MW; j++ {
			if monster[i][j] == '#' {
				monsterCoords[im] = make([]int, 2)
				monsterCoords[im][0] = i
				monsterCoords[im][1] = j
				im++
			}
		}
	}

	// Finally search for the monsters
	roughness := 0
	Rg := [2]int{1, 0}
	for i := 0; i < 8; i++ {
		if i == 4 {
			for j := 0; j < D/2; j++ {
				k := D - j - 1
				compMap[j], compMap[k] = compMap[k], compMap[j]
			}
		}
		m := rotate(compMap, Rg)

		count := 0
		for r := 0; r < D-MH; r++ {
			for c := 0; c < D-MW; c++ {
				found := true
				for im = 0; im < nHashes; im++ {
					mr := monsterCoords[im][0]
					mc := monsterCoords[im][1]
					if m[r+mr][c+mc] != '#' {
						found = false
						break
					}
				}
				if found {
					count++
					for im = 0; im < nHashes; im++ {
						mr := monsterCoords[im][0]
						mc := monsterCoords[im][1]
						m[r+mr][c+mc] = 'O'
					}
				}
			}
		}
		if count > 0 {
			roughness = 0
			for r := 0; r < D; r++ {
				for c := 0; c < D; c++ {
					if m[r][c] == '#' {
						roughness++
					}
				}
			}
			break
		}
		Rg[0], Rg[1] = -Rg[1], Rg[0] // spin by 90 deg
	}

	return strconv.Itoa(roughness)
}

func rotate(m []string, R [2]int) [][]byte {
	re := R[0]
	im := R[1]
	N := len(m)
	D := N - 1

	d := make([][]byte, N)
	for i := 0; i < N; i++ {
		d[i] = make([]byte, N)
	}
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			xn, yn := 2*x-D, 2*y-D
			xn, yn = xn*re-yn*im, xn*im+yn*re
			xn, yn = (xn+D)/2, (yn+D)/2
			d[xn][yn] = m[x][y]
		}
	}
	return d
}

func getCol(t []string, c int) string {
	N := len(t)
	d := make([]byte, N)
	for i := 0; i < N; i++ {
		d[i] = t[i][c]
	}
	return string(d)
}

func reverse(s string) string {
	N := len(s)
	L := N - 1
	d := make([]byte, N)
	for i := 0; i < N; i++ {
		d[L-i] = s[i]
	}
	return string(d)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
