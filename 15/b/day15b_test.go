package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"test1.txt", "0", "175594"},
		{"test2.txt", "0", "2578"},
		{"test3.txt", "0", "3544142"},
		{"test4.txt", "0", "261214"},
		{"test5.txt", "0", "6895259"},
		{"test6.txt", "0", "18"},
		{"test7.txt", "0", "362"},
	}

	Ntests := len(testCases)
	for i := 0; i < Ntests; i++ {
		ans := doit(testCases[i][0], testCases[i][1])
		exp := testCases[i][2]
		if ans != exp {
			t.Errorf("Result '%s' does not match expected value of '%s'", ans, exp)
		}
	}
}
