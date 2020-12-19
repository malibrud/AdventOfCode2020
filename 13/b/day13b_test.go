package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"test1.txt", "0", "1068781"},
		{"test2.txt", "0", "754018"},
		{"test3.txt", "0", "779210"},
		{"test4.txt", "0", "1261476"},
		{"test5.txt", "0", "1202161486"},
	}

	Ntests := len(testCases)
	for i := 0; i < Ntests; i++ {
		ans := doit(testCases[i][0], testCases[i][1])
		exp := testCases[i][2]
		if ans != exp {
			t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
		}
	}
}
