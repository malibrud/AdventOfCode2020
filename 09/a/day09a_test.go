package main

import (
	"testing"
)

func Test(t *testing.T) {
	ans := doit("test.txt", 5)
	exp := "127"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}
