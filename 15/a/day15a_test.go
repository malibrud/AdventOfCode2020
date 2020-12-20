package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := [][]string{
		{"test1.txt", "0", "436"},
		{"test2.txt", "0", "1"},
		{"test3.txt", "0", "10"},
		{"test4.txt", "0", "27"},
		{"test5.txt", "0", "78"},
		{"test6.txt", "0", "438"},
		{"test7.txt", "0", "1836"},
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
