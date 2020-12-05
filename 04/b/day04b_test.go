package main

import (
	"testing"
)

func Test(t *testing.T) {
	ans := doit("test.txt")
	exp := "2"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}

func TestInvalid(t *testing.T) {
	ans := doit("testInvalid.txt")
	exp := "0"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}

func TestValid(t *testing.T) {
	ans := doit("testValid.txt")
	exp := "4"
	if ans != exp {
		t.Errorf("Result '%s'does not match expected value of '%s'", ans, exp)
	}
}
