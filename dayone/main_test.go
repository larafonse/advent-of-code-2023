package main

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	line := "veightone1"
	exp := 11
	sum := getSum(line)
	if sum != exp {
		t.Fatalf("fail")
	}
}

func TestParseLine(t *testing.T) {
	line := "testonetwodeight"
	exp := "test12d"
	parsedLine := parseLine(line)
	if parsedLine != exp {
		t.Fatalf("fail")
	}
}
