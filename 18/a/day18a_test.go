package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"test1.txt", "0", "71"},
		{"test2.txt", "0", "51"},
		{"test3.txt", "0", "26"},
		{"test4.txt", "0", "437"},
		{"test5.txt", "0", "12240"},
		{"test6.txt", "0", "13632"},
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
