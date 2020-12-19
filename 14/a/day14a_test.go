package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"test1.txt", "0", "165"},
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
