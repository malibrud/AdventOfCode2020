package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type tileTy struct {
	id   int
	data []string
}

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	rawTiles := strings.Split(string(dat), "\r\n\r\n")

	fmt.Printf("Number of tiles found: %d\n", len(rawTiles))
	edgeCounts := make(map[string]int)
	tileIds := make(map[string]int)

	// parse the tiles
	for _, rawTile := range rawTiles {
		lines := strings.Split(rawTile, "\r\n")
		idstr := lines[0][5 : len(lines[0])-1]
		id, _ := strconv.Atoi(idstr)
		tile := lines[1:len(lines)]
		H := len(tile)
		W := len(tile[0])
		t := new(tileTy)
		t.id = id
		t.data = tile
		top := string(tile[0])
		bot := string(tile[H-1])
		left := getCol(tile, 0)
		right := getCol(tile, W-1)
		topR := reverse(top)
		botR := reverse(bot)
		leftR := reverse(left)
		rightR := reverse(right)
		edgeCounts[top]++
		edgeCounts[bot]++
		edgeCounts[left]++
		edgeCounts[right]++
		edgeCounts[topR]++
		edgeCounts[botR]++
		edgeCounts[leftR]++
		edgeCounts[rightR]++
		tileIds[top] = id
		tileIds[bot] = id
		tileIds[left] = id
		tileIds[right] = id
		tileIds[topR] = id
		tileIds[botR] = id
		tileIds[leftR] = id
		tileIds[rightR] = id
	}

	cornerTiles := make(map[int]int)
	for k, v := range edgeCounts {
		if v == 1 {
			id := tileIds[k]
			cornerTiles[id]++
		}
	}

	prod := int64(1)
	for id, count := range cornerTiles {
		if count == 4 {
			prod *= int64(id)
		}
	}

	return strconv.FormatInt(prod, 10)
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
