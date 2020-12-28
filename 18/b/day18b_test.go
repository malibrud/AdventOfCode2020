package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"test2.txt", "0", "51"},
		{"test6.txt", "0", "23340"},
		{"test1.txt", "0", "231"},
		{"test3.txt", "0", "46"},
		{"test4.txt", "0", "1445"},
		{"test5.txt", "0", "669060"},
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
