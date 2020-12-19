package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doit("data.txt", ""))
}

func doit(fileName string, arg string) string {
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	lines := strings.Split(string(dat), "\r\n")
	mem := make(map[int64]int64)

	mask0 := int64(0)
	mask1 := int64(0)
	//maskX := int64(0)
	nXbits := 0
	nAnyMasks := 0
	var anyMask [1024]int64 // allow room for 10 'X' bits
	re := regexp.MustCompile(`mem\[(\d+)\]`)
	for _, line := range lines {
		instruct := strings.Split(line, " = ")
		lval := instruct[0]
		rval := instruct[1]
		if lval == "mask" {
			var maskstr string

			// Mask of don't care values 0
			maskstr = rval
			maskstr = strings.ReplaceAll(maskstr, "1", "X") // should now only have 0, X
			maskstr = strings.ReplaceAll(maskstr, "0", "1") // should now only have 1, X
			maskstr = strings.ReplaceAll(maskstr, "X", "0") // should now only have 0, 1
			mask0, _ = strconv.ParseInt(maskstr, 2, 64)

			// Mask of set values 1
			maskstr = rval
			maskstr = strings.ReplaceAll(maskstr, "X", "0") // should now only have 0, 1
			mask1, _ = strconv.ParseInt(maskstr, 2, 64)

			// Masks for all permutations of the 'any' bits.
			nXbits = strings.Count(rval, "X")
			nAnyMasks = 1
			for i := 0; i < 36; i++ {
				b := 36 - 1 - i
				if rval[b] == 'X' {
					xBit := int64(1) << i
					for k := 0; k < nAnyMasks; k++ {
						anyMask[nAnyMasks+k] = anyMask[k] | xBit
					}
					nAnyMasks *= 2
				}
			}
			println("nXbits = ", nXbits)
		} else if lval[:3] == "mem" {
			fields := re.FindStringSubmatch(lval)
			addr, _ := strconv.ParseInt(fields[1], 10, 64)
			val, _ := strconv.ParseInt(rval, 10, 64)

			// Do the masking thing
			for i := 0; i < nAnyMasks; i++ {
				thisaddr := addr&mask0 | mask1 | anyMask[i]
				mem[thisaddr] = val
			}
		}
	}
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}

	return strconv.FormatInt(sum, 10)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
