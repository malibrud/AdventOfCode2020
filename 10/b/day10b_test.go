package main

import (
	"testing"
)

func Test1(t *testing.T) {
	ans := doit("test1.txt")
	exp := "8"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}

func Test2(t *testing.T) {
	ans := doit("test2.txt")
	exp := "19208"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}
