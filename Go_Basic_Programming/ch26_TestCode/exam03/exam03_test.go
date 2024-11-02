package main

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	n, err := Atoi("1")
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}
}
