package main

import (
	"bufio"
	"strings"
	"testing"
)

const example string = `3   4
4   3
2   5
1   3
3   9
3   3`

const expectedResultPart1 int = 11
const expectedResultPart2 int = 31

func TestPart1(t *testing.T) {
	input := strings.NewReader(example)
	res := Part1(bufio.NewScanner(input))
	if res != expectedResultPart1 {
		t.Errorf("got %d, expected %d", res, expectedResultPart1)
	}
}

func TestPart2(t *testing.T) {
	input := strings.NewReader(example)
	res := Part2(bufio.NewScanner(input))
	if res != expectedResultPart2 {
		t.Errorf("got %d, expected %d", res, expectedResultPart2)
	}
}
